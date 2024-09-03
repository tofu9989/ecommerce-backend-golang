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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.CreatedAt = time.Now()

	collection := config.GetCollection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	collection := config.GetCollection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, _ := collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		cur.Decode(&user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user models.User
	collection := config.GetCollection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
