package requests

import (
	"golang.org/x/exp/slices"
	"net/http"
	"testing"
	"varaapoyta-backend-refactor/responseStructures"
)

func TestGetRestaurantRequestHandler(t *testing.T) {
	expectedHeaderKeys := []string{"Content-Type", "Client_id", "User-Agent"}
	expectedHeaderValues := []string{"application/json", "jNAWMvWD9rp637RaR", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"}

	req := GetRequestHandlerFor(&Api{Name: "restaurant"})

	for key, value := range req.Header {
		if !slices.Contains(expectedHeaderKeys, key) {
			t.Errorf("GetRequestHandler - header key %v not found", key)
		}
		if !slices.Contains(expectedHeaderValues, value[0]) {
			t.Errorf("GetRequestHandler - header value %v not found", value)
		}
	}
	if req.Method != "POST" {
		t.Errorf("GetRequestHandler - expected method to be POST, got %v", req.Method)
	}
	if req.URL.String() != "https://api.raflaamo.fi/query" {
		t.Errorf("GetRequestHandler - expected url to be https://api.raflaamo.fi/query, got %v", req.URL.String())
	}
}
func TestGetGraphApiRequestHandler(t *testing.T) {
	expectedHeaderKeys := []string{"User-Agent"}
	expectedHeaderValues := []string{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"}

	req := GetRequestHandlerFor(&Api{Name: "graph", Url: "https://www.google.com"})

	for key, value := range req.Header {
		if !slices.Contains(expectedHeaderKeys, key) {
			t.Errorf("GetRequestHandler - header key %v not found", key)
		}
		if !slices.Contains(expectedHeaderValues, value[0]) {
			t.Errorf("GetRequestHandler - header value %v not found", value)
		}
	}
	if req.Method != "GET" {
		t.Errorf("GetRequestHandler - expected method to be POST, got %v", req.Method)
	}
	if req.URL.String() != "https://www.google.com" {
		t.Errorf("GetRequestHandler - expected url to be https://www.google.com, got %v", req.URL.String())
	}
}

func TestIsGraphApi(t *testing.T) {
	if !isGraphApi(&Api{Name: "graph"}) {
		t.Errorf("isGraphApi - expected true, got false")
	}
	if isGraphApi(&Api{Name: "restaurant"}) {
		t.Errorf("isGraphApi - expected false, got true")
	}
}

func TestIsRestaurantsApi(t *testing.T) {
	if !isRestaurantsApi(&Api{Name: "restaurant"}) {
		t.Errorf("isRestaurantsApi - expected true, got false")
	}
	if isRestaurantsApi(&Api{Name: "graph"}) {
		t.Errorf("isRestaurantsApi - expected false, got true")
	}
}
func TestGraphApiUrlExists(t *testing.T) {
	if !graphApiUrlExists(&Api{Url: "https://www.google.com"}) {
		t.Errorf("graphApiUrlExists - expected true, got false")
	}
	if graphApiUrlExists(&Api{Url: ""}) {
		t.Errorf("graphApiUrlExists - expected false, got true")
	}
}

func TestSetGraphApiHeadersTo(t *testing.T) {
	expectedHeaderKeys := []string{"User-Agent"}
	expectedHeaderValues := []string{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"}

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	setGraphApiHeadersTo(req)

	for key, value := range req.Header {
		if !slices.Contains(expectedHeaderKeys, key) {
			t.Errorf("setGraphApiHeadersTo - header key %v not found", key)
		}
		if !slices.Contains(expectedHeaderValues, value[0]) {
			t.Errorf("setGraphApiHeadersTo - header value %v not found", value)
		}
	}
}

func TestSetRestaurantApiHeadersTo(t *testing.T) {
	expectedHeaderKeys := []string{"Content-Type", "Client_id", "User-Agent"}
	expectedHeaderValues := []string{"application/json", "jNAWMvWD9rp637RaR", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"}

	req, _ := http.NewRequest("POST", "https://api.raflaamo.fi/query", nil)
	setRestaurantApiHeadersTo(req)

	for key, value := range req.Header {
		if !slices.Contains(expectedHeaderKeys, key) {
			t.Errorf("setRestaurantApiHeadersTo - header key %v not found", key)
		}
		if !slices.Contains(expectedHeaderValues, value[0]) {
			t.Errorf("setRestaurantApiHeadersTo - header value %v not found", value)
		}
	}
}

func TestDeserializeRestaurantApiResponse(t *testing.T) {
	response := `{"data": {"listRestaurantsByLocation": {"totalCount": 470}}}`
	expectedTotalCount := 470

	responseStruct := responseStructures.RestaurantApiResponse{}

	bytes := []byte(response)
	restaurant, err := deserializeResponse(bytes, responseStruct)
	if err != nil {
		t.Errorf("TestDeserializeRestaurantApiResponse - expected no error, got %v", err)
	}

	restaurantWithType, ok := restaurant.(*responseStructures.RestaurantApiResponse)
	if !ok {
		t.Errorf("TestDeserializeRestaurantApiResponse - expected restaurant to be of type RestaurantApiResponse, got %T", restaurant)
	}
	if restaurantWithType.Data.ListRestaurantsByLocation.TotalCount != expectedTotalCount {
		t.Errorf("TestDeserializeRestaurantApiResponse - expected total count to be %v, got %v", expectedTotalCount, restaurantWithType.Data.ListRestaurantsByLocation.TotalCount)
	}
}

func TestDeserializeGraphApiResponse(t *testing.T) {
	response := `[{"name": "Stone's", "intervals": [{"from": 1660322700000,"to": 1660322700000,"color": "transparent"}]}]`
	expectedName := "Stone's"

	responseStruct := responseStructures.GraphApiResponse{}

	bytes := []byte(response)
	restaurant, err := deserializeResponse(bytes, responseStruct)
	if err != nil {
		t.Errorf("TestDeserializeGraphApiResponse - expected no error, got %v", err)
	}

	restaurantWithType, ok := restaurant.(*responseStructures.GraphApiResponse)
	if !ok {
		t.Errorf("TestDeserializeGraphApiResponse - expected restaurant to be of type GraphApiResponse, got %T", restaurant)
	}
	if (*restaurantWithType)[0].Name != expectedName {
		t.Errorf("TestDeserializeGraphApiResponse - expected name to be %v, got %v", expectedName, (*restaurantWithType)[0].Name)
	}
}
