package requests

import (
	"bytes"
	"log"
	"net/http"
)

type Api struct {
	Name string
	Url  string
}

func GetRequestHandlerFor(api *Api) *http.Request {
	if graphApi(api) {
		if !graphApiUrlExists(api) {
			log.Fatal("[GetRequestHandlerFor] - API Name was graph but url was empty.")
		}
		req, err := http.NewRequest("GET", api.Url, nil)

		if err != nil {
			log.Fatal("There was an error initializing the graph api handler.")
		}
		setGraphApiHeadersTo(req)
		return req
	} else if restaurantsApi(api) {
		payload := getRestaurantsApiPayload()
		req, err := http.NewRequest("POST", "https://api.raflaamo.fi/query", bytes.NewBuffer(payload))
		if err != nil {
			log.Fatalln("Unexpected error when we constructed the restaurant api handler.")
		}
		setRestaurantApiHeadersTo(req)
		return req
	} else {
		log.Fatal("[GetRequestHandlerFor] - Invalid api name")
	}
	return nil
}
func graphApi(api *Api) bool {
	return api.Name == "graph"
}
func graphApiUrlExists(api *Api) bool {
	if api.Url == "" {
		return false
	}
	return true
}
func setGraphApiHeadersTo(r *http.Request) {
	r.Header.Add("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
}

func restaurantsApi(api *Api) bool {
	return api.Name == "restaurant"
}

func getRestaurantsApiPayload() []byte {
	const payload = `{"operationName":"getRestaurantsByLocation","variables":{"first":1000,"input":{"restaurantType":"ALL","locationName":"Helsinki","feature":{"rentableVenues":false}},"after":"eyJmIjowLCJnIjp7ImEiOjYwLjE3MTE2LCJvIjoyNC45MzI1OH19"},"query":"fragment Locales on LocalizedString {fi_FI }fragment Restaurant on Restaurant {  id  name {    ...Locales    }  address {    municipality {      ...Locales       }        street {      ...Locales       }       zipCode     }   openingTime {    restaurantTime {      ranges {        start        end             }             }    kitchenTime {      ranges {        start        end        endNextDay              }             }    }  links {    tableReservationLocalized {      ...Locales        }    homepageLocalized {      ...Locales          }   }     }query getRestaurantsByLocation($first: Int, $after: String, $input: ListRestaurantsByLocationInput!) {  listRestaurantsByLocation(first: $first, after: $after, input: $input) {    totalCount      edges {      ...Restaurant        }     }}"}`
	payloadBytes := []byte(payload)
	return payloadBytes
}

func setRestaurantApiHeadersTo(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client_id", "jNAWMvWD9rp637RaR")
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
}
