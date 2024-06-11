package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var(
	doctorCollection *mongo.Collection
	patientCollection *mongo.Collection
)

func init() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file: %v", err)
	}

	connectionString := os.Getenv("MONGODB_URL")

	if connectionString == ""{
		log.Fatalf("MONGODB_URL not sen in .env file")
	}

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	//Check the Connection
	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Successful")

	//setup collection
	doctorCollection = client.Database("Hospital").Collection("Doctor")
	patientCollection = client.Database("Hospital").Collection("Patient")
}

//get doctor db coll
func GetDoctorCollection() *mongo.Collection{
	return doctorCollection
}

// GetPatientCollection returns the patient collection
func GetPatientCollection() *mongo.Collection {
	return patientCollection
}