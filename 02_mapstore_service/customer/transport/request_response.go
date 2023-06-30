package transport

import (
	"context"
	"encoding/json"
	"gokit-microservices/02_mapstore_service/customer"
	"net/http"
)

type createRequest struct {
	Customer customer.Customer `json:"customer"`
}

type createResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type getByIDRequest struct {
	ID string `json:"id"`
}

type getByIDResponse struct {
	Customer customer.Customer `json:"customer"`
	Error    error             `json:"error"`
}

type deleteRequest struct {
	ID string `json:"id"`
}

type deleteResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// generic response encoding function
func encodeRespose(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
