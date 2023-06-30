package http

import (
	"context"
	"encoding/json"
	"errors"
	"gokit-microservices/02_mapstore_service/customer/transport"

	"net/http"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var ErrRouting = errors.New("bad routing")

// wiring Go kit endpoints to the HTTP transport
func NewService(endpoints transport.Endpoints, logger log.Logger) http.Handler {

	var (
		r = mux.NewRouter()
		//errorLogger = khttp.ServerErrorLogger(logger)
		//errorEncoder = khttp.ServerErrorEncoder(nil)
	)
	//options = append(options, errorLogger)

	r.Methods("POST").Path("/customer/create").Handler(khttp.NewServer(endpoints.Create, decodeCreateRequest, encodeRespose))
	r.Methods("GET").Path("/customer/{id}").Handler(khttp.NewServer(endpoints.GetByID, decodeGetByIDRequest, encodeRespose))
	r.Methods("DELETE").Path("customer/{id}").Handler(khttp.NewServer(endpoints.Delete, decodeDeleteRequest, encodeRespose))

	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.GetByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// generic response encoding function
func encodeRespose(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
