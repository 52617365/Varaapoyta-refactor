package requests

import (
	"golang.org/x/exp/slices"
	"testing"
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
