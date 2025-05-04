package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Program struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	FacultyID primitive.ObjectID `bson:"faculty_id"` // Связь 1:M
	Duration  int                `bson:"duration_years"`
}
