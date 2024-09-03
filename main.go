package main

import (
	"ecommerce-backend-golang/config"
	"ecommerce-backend-golang/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Connect to MongoDB
	config.ConnectDB()

	// Initialize routes
	routes.ProductRoutes(router)
	routes.OrderRoutes(router)
	routes.UserRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
