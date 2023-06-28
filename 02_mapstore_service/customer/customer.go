package customer

import "context"

type Customer struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Address []Address `json:"address"`
}

type Address struct {
	Type    string `json:"type"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Line3   string `json:"line3"`
	Line4   string `json:"line4"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Repository interface {
	CreateCustomer(context.Context, Customer) error
	GetCustomerByID(context.Context, string) (Customer, error)
	DeleteCustomer(context.Context, string) (string, error)
}
