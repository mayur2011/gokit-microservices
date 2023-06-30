package implementation

import (
	"context"
	"gokit-microservices/02_mapstore_service/customer"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type customerService struct {
	repository customer.Repository
	logger     log.Logger
}

func NewCustomerService(repo customer.Repository,
	logger log.Logger) customer.CustomerService {
	return &customerService{
		repository: repo,
		logger:     logger,
	}
}

func (svc *customerService) Create(ctx context.Context, cust customer.Customer) (string, error) {
	logger := log.With(svc.logger, "method", "Create")

	if err := svc.repository.CreateCustomer(ctx, cust); err != nil {
		level.Error(logger).Log("err", err)
		return "Failed", err
	}
	return "Success", nil
}

func (svc *customerService) GetByID(ctx context.Context, id string) (customer.Customer, error) {
	var customerDoc customer.Customer

	return customerDoc, nil
}

func (svc *customerService) Delete(ctx context.Context, id string) (string, error) {

	return "", nil
}
