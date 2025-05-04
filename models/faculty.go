package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Faculty struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Building string             `bson:"building"`
}
