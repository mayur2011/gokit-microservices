package transport

import (
	"context"
	"encoding/json"
	"net/http"
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

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
