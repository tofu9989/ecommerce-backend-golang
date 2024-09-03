package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"ecommerce-backend-golang/config"
	"ecommerce-backend-golang/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	collection := config.GetCollection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, product)
	json.NewEncoder(w).Encode(result)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	collection := config.GetCollection("products")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, _ := collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var product models.Product
		cur.Decode(&product)
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
