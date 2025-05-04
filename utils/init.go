package utils

import (
	"RBD_dev/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ApplicantCollection   *mongo.Collection
	DocumentCollection    *mongo.Collection
	FacultyCollection     *mongo.Collection
	ProgramCollection     *mongo.Collection
	ApplicationCollection *mongo.Collection
)

func InitCollections() {
	db := config.Client.Database("admission")

	ApplicantCollection = db.Collection("applicants")
	DocumentCollection = db.Collection("documents")
	FacultyCollection = db.Collection("faculties")
	ProgramCollection = db.Collection("programs")
	ApplicationCollection = db.Collection("applications")
}
