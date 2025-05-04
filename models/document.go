package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ApplicantID primitive.ObjectID `bson:"applicant_id"` // Связь 1:1
	Passport    string             `bson:"passport"`
	SNILS       string             `bson:"snils"`
}
