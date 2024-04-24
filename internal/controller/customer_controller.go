package controller

import (
	"encoding/json"
	helper "go-service-boooking-to-go/internal/helpers"
	"go-service-boooking-to-go/internal/usecase"
	"net/http"
)

type CustomerController struct {
	customerUseCase *usecase.CustomerUseCase
}

func NewCustomerController(customerUseCase *usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUseCase}
}

func (customerController *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers, err := customerController.customerUseCase.GetCustomer()
	if err != nil {
		webResponse := helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "error",
			Data:   nil,
			// Message: "Failed to get customer",
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(webResponse)
		return
	}
	webResponse := helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Data:    customers,
		Message: "Successfull get data customer",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}
