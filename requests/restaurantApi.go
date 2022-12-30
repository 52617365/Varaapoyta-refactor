package requests

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"varaapoyta-backend-refactor/responseStructures"
	"varaapoyta-backend-refactor/time"
)

type Restaurants struct {
	restaurantWithTimeSlots *RestaurantWithTimeSlots
	err                     error
}

type RestaurantWithTimeSlots struct {
	Restaurant               *responseStructures.Edges `json:"restaurant"`
	TimeSlots                []string                  `json:"timeSlots"`
	TimeTillRestaurantCloses *time.RelativeTime        `json:"timeTillRestaurantCloses"`
	TimeTillKitchenCloses    *time.RelativeTime        `json:"timeTillKitchenCloses"`
}

func GetRestaurantsWithTimeSlots(city string) ([]RestaurantWithTimeSlots, error) {
	restaurants, err := GetRestaurants(city)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	restaurantsWithTimeSlots := make(chan Restaurants, len(restaurants))

	for _, restaurant := range restaurants {
		wg.Add(1)
		restaurant := restaurant
		go func() {
			defer wg.Done()
			timeSlots, err := GetGraphApiTimeSlotsFrom(restaurant.ReservationPageID)
			if err != nil {
				restaurantsWithTimeSlots <- Restaurants{restaurantWithTimeSlots: nil, err: err}
				return
			}
			if requiredInfoExists(timeSlots, restaurant.OpeningTime.KitchenTime.Ranges) {
				timeSlots = time.ExtractUnwantedTimeSlots(timeSlots, getKitchenClosingTime(&restaurant))
				timeToRestaurantClosing := time.CalcRelativeTimeToFromCurrentTime(getRestaurantClosingTime(&restaurant))
				timeToKitchenClosing := time.CalcRelativeTimeToFromCurrentTime(getKitchenClosingTime(&restaurant))
				restaurantWithTimeSlots := RestaurantWithTimeSlots{Restaurant: &restaurant, TimeSlots: timeSlots, TimeTillRestaurantCloses: timeToRestaurantClosing, TimeTillKitchenCloses: timeToKitchenClosing}
				restaurantsWithTimeSlots <- Restaurants{restaurantWithTimeSlots: &restaurantWithTimeSlots, err: nil}
			} else {
				restaurantsWithTimeSlots <- Restaurants{restaurantWithTimeSlots: nil, err: &RequiredInfoDoesNotExist{}}
			}
		}()
	}

	wg.Wait()
	close(restaurantsWithTimeSlots)

	syncedRestaurantsWithTimeSlots, err := syncRestaurantsWithTimeSlots(restaurantsWithTimeSlots)
	if err != nil {
		return nil, err
	}
	return syncedRestaurantsWithTimeSlots, nil
}

func requiredInfoExists(timeSlots []string, kitchenClosingRanges []responseStructures.Ranges) bool {
	if kitchenClosingRanges == nil || len(timeSlots) == 0 || kitchenClosingRanges[0].End == "" {
		return false
	}
	return true
}

func getKitchenClosingTime(restaurant *responseStructures.Edges) string {
	if restaurant.OpeningTime.KitchenTime.Ranges == nil {
		log.Fatal("getKitchenClosingTime - restaurant.OpeningTime.KitchenTime.Ranges is nil")
	}
	kitchenClosingTime := restaurant.OpeningTime.KitchenTime.Ranges[0].End

	return kitchenClosingTime
}
func getRestaurantClosingTime(restaurant *responseStructures.Edges) string {
	if restaurant.OpeningTime.RestaurantTime.Ranges == nil {
		log.Fatal("getRestaurantClosingTime - restaurant.OpeningTime.RestaurantTime.Ranges is nil")
	}
	restaurantClosingTime := restaurant.OpeningTime.RestaurantTime.Ranges[0].End

	return restaurantClosingTime
}

func syncRestaurantsWithTimeSlots(restaurantsWithTimeSlots chan Restaurants) ([]RestaurantWithTimeSlots, error) {
	syncedrestaurantsWithTimeSlots := make([]RestaurantWithTimeSlots, 0, len(restaurantsWithTimeSlots))
	for restaurantWithTimeSlot := range restaurantsWithTimeSlots {
		if restaurantWithTimeSlot.err != nil {
			if errors.Is(restaurantWithTimeSlot.err, &RequiredInfoDoesNotExist{}) {
				continue
			}
			return nil, restaurantWithTimeSlot.err
		}
		if !timeSlotsFound(restaurantWithTimeSlot.restaurantWithTimeSlots.TimeSlots) {
			continue
		}
		syncedrestaurantsWithTimeSlots = append(syncedrestaurantsWithTimeSlots, *restaurantWithTimeSlot.restaurantWithTimeSlots)
	}
	return syncedrestaurantsWithTimeSlots, nil
}

func timeSlotsFound(timeSlots []string) bool {
	return len(timeSlots) != 0
}

func GetRestaurants(city string) ([]responseStructures.Edges, error) {
	api := &Api{
		Name: "restaurant",
	}
	req := GetRequestHandlerFor(api)

	resp, err := getRestaurantsFromApi(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("GetRestaurants - could not get restaurants from api.raflaamo.fi/query")
	}
	restaurants, err := ReadRestaurantApiResponse(resp)

	if err != nil {
		return nil, err
	}
	validRestaurants := filterRestaurants(restaurants, city)
	return validRestaurants, nil
}

func getRestaurantsFromApi(req *http.Request) (*http.Response, error) {
	resp, err := sendRequest(req)
	return resp, err
}

var ReadRestaurantApiResponse = func(res *http.Response) (*responseStructures.RestaurantApiResponse, error) {
	readBuffer, err := ReadResponseBuffer(res)
	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	deserializedResponse, err := deserializeRestaurantApiResponse(readBuffer)
	if err != nil {
		return &responseStructures.RestaurantApiResponse{}, err
	}
	return deserializedResponse, nil
}

func deserializeRestaurantApiResponse(response []byte) (*responseStructures.RestaurantApiResponse, error) {
	deserializedResponse, err := deserializeResponse(response, &responseStructures.RestaurantApiResponse{})
	if err != nil {
		return nil, err
	}
	result := deserializedResponse.(*responseStructures.RestaurantApiResponse)
	return result, nil
}

func filterRestaurants(restaurants *responseStructures.RestaurantApiResponse, city string) []responseStructures.Edges {
	validRestaurants := getValidRestaurants(restaurants, city)
	setReservationIdsToRestaurants(validRestaurants)
	return validRestaurants
}

func getValidRestaurants(restaurants *responseStructures.RestaurantApiResponse, city string) []responseStructures.Edges {
	validRestaurants := make([]responseStructures.Edges, 0, 50)
	for _, restaurant := range restaurants.Data.ListRestaurantsByLocation.Edges {
		if reservationPageExists(restaurant.Links.TableReservationLocalized.FiFI) && restaurantIsFromCorrectCity(restaurant.Address.Municipality.FiFI, city) {
			validRestaurants = append(validRestaurants, restaurant)
		}
	}
	return validRestaurants
}

// setReservationIdsToRestaurants the id returned from the endpoint is not the same as the one in the reservation page url.
// we need the latter to access the graph api.
func setReservationIdsToRestaurants(restaurants []responseStructures.Edges) {
	for index := range restaurants {
		reservationPageUrl := restaurants[index].Links.TableReservationLocalized.FiFI
		restaurantId, err := getReservationIdFrom(reservationPageUrl)
		if err != nil {
			// TODO: this should be logged because it should not have a problem finding the id after the check.
			continue
		}
		restaurants[index].ReservationPageID = restaurantId
	}
}

var RegexToMatchRestaurantId = regexp.MustCompile(`fi/(\d+)`)

func getReservationIdFrom(reservationPageUrl string) (string, error) {
	restaurantIdMatch := RegexToMatchRestaurantId.FindAllStringSubmatch(reservationPageUrl, -1)
	if restaurantIdMatch == nil {
		return "", errors.New("getReservationIdFrom - could not get restaurant id from reservation page url")
	}
	restaurantId := restaurantIdMatch[0][1]
	return restaurantId, nil
}

func reservationPageExists(reservationUrl string) bool {
	if reservationUrl == "" {
		return false
	}
	if !strings.Contains(reservationUrl, "https://s-varaukset.fi/online/reservation/fi/") {
		return false
	}
	return true
}

func restaurantIsFromCorrectCity(restaurantCity string, usersCity string) bool {
	return strings.EqualFold(restaurantCity, usersCity)
}
