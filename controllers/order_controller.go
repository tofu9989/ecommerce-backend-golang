package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"ecommerce-backend-golang/config"
	"ecommerce-backend-golang/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.CreatedAt = time.Now()

	collection := config.GetCollection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, order)
	json.NewEncoder(w).Encode(result)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Order
	collection := config.GetCollection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, _ := collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var order models.Order
		cur.Decode(&order)
		orders = append(orders, order)
	}
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var order models.Order
	collection := config.GetCollection("orders")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}
