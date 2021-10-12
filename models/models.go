package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Complete string             `bson:"complete"`
	Author   string             `bson:"author"`
}
