package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Applicant struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	BirthDate string             `bson:"birth_date"`
}
