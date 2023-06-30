package customer

import (
	"context"
	"errors"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCmdRepository    = errors.New("unable to command repository")
	ErrQueryRepository  = errors.New("unable to query repository")
)

type CustomerService interface {
	Create(context.Context, Customer) (string, error)
	GetByID(context.Context, string) (Customer, error)
	Delete(context.Context, string) (string, error)
}
