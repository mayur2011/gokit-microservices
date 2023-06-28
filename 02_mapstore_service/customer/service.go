package customer

type CustomerService interface {
	Create(Customer) (string, error)
	GetByID(string) (Customer, error)
	Delete(string) (string, error)
}
