package routes

import (
	"ecommerce-backend-golang/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", controllers.GetOrderByID).Methods("GET")
}
