package transport

import (
	"context"
	"gokit-microservices/00_string_service_basic_3/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Upperstring endpoint.Endpoint
	Countstring endpoint.Endpoint
}

func NewEndpoints(srvc service.StringService) Endpoints {
	return Endpoints{
		Upperstring: makeUppercaseEndpoint(srvc),
		Countstring: makeCountEndpoint(srvc),
	}
}

func makeUppercaseEndpoint(srvc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		output, err := srvc.Uppercase(req.Input)
		if err != nil {
			return uppercaseResponse{output, err.Error()}, err
		}
		return uppercaseResponse{output, ""}, nil
	}
}

func makeCountEndpoint(srvc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		output := srvc.Count(req.Input)
		return countResponse{output}, nil
	}
}
