package mapstoredb

import (
	"context"
	"errors"
	"fmt"
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
		return errors.New("customer exist for ID =" + cust.ID)
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
	return customerDoc, customer.ErrCustomerNotFound
}

func (repo *repository) DeleteCustomer(_ context.Context, id string) error {
	if repo.isCustomerExist(id) {
		delete(repo.store, id)
		return nil
	}
	return fmt.Errorf("customer not found")
}

//local function to check customer exists
func (repo *repository) isCustomerExist(id string) bool {
	_, ok := repo.store[id]
	return ok
}
