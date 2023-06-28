package implementation

import (
	customerdmn "gokit-microservices/02_mapstore_service/customer"

	"github.com/go-kit/log"
)

type customerService struct {
	repository customerdmn.Repository
	logger     log.Logger
}

func (cust customerService) Create(customerdmn.Customer) (string, error) {

	return "", nil
}
