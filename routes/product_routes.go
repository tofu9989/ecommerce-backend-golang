package routes

import (
	"ecommerce-backend-golang/controllers"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
}
