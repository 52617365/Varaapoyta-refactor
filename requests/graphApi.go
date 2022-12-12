package requests

import (
	"fmt"
	"net/http"
	"varaapoyta-backend-refactor/responseStructures"
)

//	func GetGraphApiTimeSlotsFrom(restaurantId string) ([]string, error) {
//		urls := GetGraphApiUrls(restaurantId)
//		var timeSlots []string
//		for _, url := range urls {
//			//response, err := GetTimeSlotFrom(url)
//			//if err != nil {
//			//	return nil, fmt.Errorf("GetTimeSlotFrom - Error getting time slot from graph api. - %w", err)
//			//}
//			// TODO: we will probably iterate the time slots here and append to the timeSlots slice ignoring duplicates.
//			futureTimeSlot := []string{}
//			timeSlots = append(timeSlots, futureTimeSlot...)
//		}
//		return timeSlots, nil
//	}
func GetTimeSlotFrom(requestUrl string) (*responseStructures.GraphApiResponse, error) { // this will be returning a string slice later on.
	response, err := getResponseFromGraphApi(requestUrl)
	if err != nil {
		return nil, err
	}
	deserializedResponse, err := deserializeGraphApiResponse(response)
	if err != nil {
		return nil, fmt.Errorf("deserializeGraphApiResponse - Error deserializing response. - %w", err)
	}
	// TODO: here we want to extract the 15 minute time intervals in between from and to and ignore duplicates
	return deserializedResponse, nil
}

func getResponseFromGraphApi(requestUrl string) ([]byte, error) {
	requestHandler := getGraphApiRequestHandler(requestUrl)
	response, err := sendRequestToGraphApi(requestHandler)
	if err != nil {
		return nil, fmt.Errorf("sendRequestToGraphApi - Error sending request to Raflaamo graph api. - %w", err)
	}
	responseBuffer, err := ReadResponseBuffer(response)
	if err != nil {
		return nil, fmt.Errorf("ReadResponseBuffer - Error reading response buffer. - %w", err)
	}
	return responseBuffer, nil
}

func getGraphApiRequestHandler(requestUrl string) *http.Request {
	graphApi := Api{
		Name: "graph",
		Url:  requestUrl,
	}
	requestHandler := GetRequestHandlerFor(&graphApi)
	return requestHandler
}

var sendRequestToGraphApi = func(requestHandler *http.Request) (*http.Response, error) {
	response, err := sendRequest(requestHandler)
	return response, err
}

func deserializeGraphApiResponse(responseBuffer []byte) (*responseStructures.GraphApiResponse, error) {
	deserializedType := responseStructures.GraphApiResponse{}
	deserializedResponse, err := deserializeResponse(responseBuffer, &deserializedType)
	if err != nil {
		return &responseStructures.GraphApiResponse{}, nil
	}
	result := deserializedResponse.(*responseStructures.GraphApiResponse)
	return result, nil
}
