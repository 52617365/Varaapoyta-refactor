package requests

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

func GetRestaurants() (string, error) {
	req := GetRequestHandler()

	resp, err := GetRestaurantsFromApi(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}
	response, err := ReadResponseBuffer(resp)
	if err != nil {
		return "", err
	}
	return response, nil
}

func GetRequestHandler() *http.Request {
	payload := getPayload()
	req, err := http.NewRequest("POST", "https://api.raflaamo.fi/query", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("GetRestaurants - Threw an unexpected error.")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("client_id", "jNAWMvWD9rp637RaR")
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	return req
}
func getPayload() []byte {
	const payload = `{"operationName":"getRestaurantsByLocation","variables":{"first":1000,"input":{"restaurantType":"ALL","locationName":"Helsinki","feature":{"rentableVenues":false}},"after":"eyJmIjowLCJnIjp7ImEiOjYwLjE3MTE2LCJvIjoyNC45MzI1OH19"},"query":"fragment Locales on LocalizedString {fi_FI }fragment Restaurant on Restaurant {  id  name {    ...Locales    }  address {    municipality {      ...Locales       }        street {      ...Locales       }       zipCode     }   openingTime {    restaurantTime {      ranges {        start        end             }             }    kitchenTime {      ranges {        start        end        endNextDay              }             }    }  links {    tableReservationLocalized {      ...Locales        }    homepageLocalized {      ...Locales          }   }     }query getRestaurantsByLocation($first: Int, $after: String, $input: ListRestaurantsByLocationInput!) {  listRestaurantsByLocation(first: $first, after: $after, input: $input) {    totalCount      edges {      ...Restaurant        }     }}"}`
	payloadBytes := []byte(payload)
	return payloadBytes
}

var GetRestaurantsFromApi = func(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.New("GetRestaurants - could not get response from api.raflaamo.fi/query")
	}
	return resp, nil
}

var ReadResponseBuffer = func(res *http.Response) (string, error) {
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
