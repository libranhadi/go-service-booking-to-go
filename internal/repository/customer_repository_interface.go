package repository

import "go-service-boooking-to-go/internal/model"

type ICustomerRepository interface {
	GetCustomer() ([]model.Customer, error)
}
