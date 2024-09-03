package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID  primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Quantity   int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	TotalPrice float64            `json:"total_price,omitempty" bson:"total_price,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
