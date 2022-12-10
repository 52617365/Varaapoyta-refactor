package requests

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

type Api struct {
	Name string
	Url  string
}

var sendRequest = func(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.New("error sending request to Raflaamo")
	}
	return resp, nil
}

func GetRequestHandlerFor(api *Api) *http.Request {
	if isGraphApi(api) {
		return getGraphApiHandler(api)
	} else if isRestaurantsApi(api) {
		return getRestaurantsApiHandler()
	} else {
		log.Fatal("[GetRequestHandlerFor] - Invalid api name")
	}
	return nil
}

func isGraphApi(api *Api) bool {
	return api.Name == "graph"
}

func getGraphApiHandler(api *Api) *http.Request {
	if !graphApiUrlExists(api) {
		log.Fatal("[GetRequestHandlerFor] - API Name was graph but url was empty.")
	}
	req, err := http.NewRequest("GET", api.Url, nil)

	if err != nil {
		log.Fatal("There was an error initializing the graph api handler.")
	}
	setGraphApiHeadersTo(req)
	return req
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

func isRestaurantsApi(api *Api) bool {
	return api.Name == "restaurant"
}
func getRestaurantsApiHandler() *http.Request {
	payload := getRestaurantsApiPayload()
	req, err := http.NewRequest("POST", "https://api.raflaamo.fi/query", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Unexpected error when we constructed the restaurant api handler.")
	}
	setRestaurantApiHeadersTo(req)
	return req
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

var ReadResponseBuffer = func(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
