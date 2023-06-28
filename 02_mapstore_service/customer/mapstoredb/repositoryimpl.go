package mapstoredb

import (
	"context"
	"errors"
	"gokit-microservices/02_mapstore_service/customer"

	"github.com/go-kit/log"
)

type repository struct {
	store  map[string]customer.Customer
	logger log.Logger
}

func NewRepository(logger log.Logger) (customer.Repository, error) {
	return &repository{
		store:  make(map[string]customer.Customer),
		logger: log.With(logger, "rep", "mapstoredb"),
	}, nil
}

func (repo *repository) CreateCustomer(_ context.Context, cust customer.Customer) error {
	if repo.isCustomerExist(cust.ID) {
		return errors.New("customer already created for given id =" + cust.ID)
	}
	repo.store[cust.ID] = cust
	return nil
}

func (repo *repository) GetCustomerByID(_ context.Context, id string) (customer.Customer, error) {
	var customerDoc = customer.Customer{}
	if repo.isCustomerExist(id) {
		customerDoc = repo.store[id]
		return customerDoc, nil
	}
	return customerDoc, errors.New("customer does not exist for given id=" + id)
}

func (repo *repository) DeleteCustomer(_ context.Context, id string) (string, error) {
	if repo.isCustomerExist(id) {
		delete(repo.store, id)
		return "Success", nil
	}
	return "Fail", errors.New("customer does not exist for given id=" + id)
}

//local function to check customer exists
func (repo *repository) isCustomerExist(id string) bool {
	_, ok := repo.store[id]
	return ok
}
