package transport

import (
	"gokit-microservices/02_mapstore_service/customer"
)

type CreateRequest struct {
	Customer customer.Customer `json:"customer"`
}

type CreateResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	Customer customer.Customer `json:"customer"`
	Error    error             `json:"error"`
}

type DeleteRequest struct {
	ID string
}

type DeleteResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
