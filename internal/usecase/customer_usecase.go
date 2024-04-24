package usecase

import (
	"go-service-boooking-to-go/internal/model"
	"go-service-boooking-to-go/internal/repository"
)

type CustomerUseCase struct {
	iCustomerRepo repository.ICustomerRepository
}

func NewCustomerUseCase(iCustomerRepo repository.ICustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{iCustomerRepo: iCustomerRepo}
}

func (uc *CustomerUseCase) GetCustomer() ([]model.Customer, error) {
	return uc.iCustomerRepo.GetCustomer()
}
