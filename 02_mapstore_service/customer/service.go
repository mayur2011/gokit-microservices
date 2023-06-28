package customer

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCmdRepository    = errors.New("unable to command repository")
	ErrQueryRepository  = errors.New("unable to query repository")
)

type CustomerService interface {
	Create(Customer) (string, error)
	GetByID(string) (Customer, error)
	Delete(string) (string, error)
}
