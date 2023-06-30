package transport

import (
	"context"
	"gokit-microservices/02_mapstore_service/customer"

	"github.com/go-kit/kit/endpoint"
)

// Endpoint will wrap Service methods
type Endpoints struct {
	Create  endpoint.Endpoint
	GetByID endpoint.Endpoint
	Delete  endpoint.Endpoint
}

func NewEndpoints(svc customer.CustomerService) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(svc),
		GetByID: makeGetByIDEndpoint(svc),
		Delete:  makeDeleteEndpoint(svc),
	}
}

// makeEndpoint func accept service as input and expose Endpoint type
func makeCreateEndpoint(svc customer.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		result, err := svc.Create(ctx, req.Customer)
		if err != nil {
			return createResponse{"", err}, err
		}
		return createResponse{result, nil}, nil
	}
}

func makeGetByIDEndpoint(svc customer.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIDRequest)
		result, err := svc.GetByID(ctx, req.ID)
		if err != nil {
			return getByIDResponse{customer.Customer{}, err}, err
		}
		return getByIDResponse{result, nil}, nil
	}
}

func makeDeleteEndpoint(svc customer.CustomerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteRequest)
		result, err := svc.Delete(ctx, req.ID)
		if err != nil {
			return deleteResponse{"", err}, err
		}
		return deleteResponse{result, nil}, nil
	}
}

/*
type Name func(int) string

func intToString(num string) Name {
	return func(n int) string {
		s := strconv.Itoa(n)
		return s + "-" + num
	}
}
*/
