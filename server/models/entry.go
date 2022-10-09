package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID          primitive.ObjectID `bson:id`
	Dish        *string            `json:"dish"`
	Fat         *float32           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *float64           `json: "calories"`
}
