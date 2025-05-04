package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Application struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ApplicantID primitive.ObjectID `bson:"applicant_id"` // 1:M
	ProgramID   primitive.ObjectID `bson:"program_id"`
	Status      string             `bson:"status"`
	SubmittedAt string             `bson:"submitted_at"` // дата подачи
}
