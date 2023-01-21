package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoStruct struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
}
