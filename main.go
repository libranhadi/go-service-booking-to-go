package main

import (
	"go-service-boooking-to-go/config"
	"go-service-boooking-to-go/internal/controller"
	"go-service-boooking-to-go/internal/repository"
	"go-service-boooking-to-go/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.InitDb()
	router := mux.NewRouter()
	customerRepository := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUseCase(customerRepository)
	controller := controller.NewCustomerController(customerUsecase)
	router.HandleFunc("/customers", controller.GetCustomer).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
