package transport

import (
	"context"
	"encoding/json"
	"gokit-microservices/00_string_service_basic_2/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	Input string `json:"input"`
}

type uppercaseResponse struct {
	Output string `json:"output"`
	Err    string `json:"err,omitempty"`
}

type countRequest struct {
	Input string `json:"input"`
}

type countResponse struct {
	Output int `json:"output"`
}

func MakeUppercaseEndpoint(srvc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		output, err := srvc.Uppercase(req.Input)
		if err != nil {
			return uppercaseResponse{output, err.Error()}, err
		}
		return uppercaseResponse{output, ""}, nil
	}
}

func MakeCountEndpoint(srvc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		output := srvc.Count(req.Input)
		return countResponse{output}, nil
	}
}

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
