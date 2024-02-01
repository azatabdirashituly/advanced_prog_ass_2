package cmd

import "go.mongodb.org/mongo-driver/bson/primitive"

// Teacher is a struct that represents a teacher
type Teacher struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName"`
	LastName  string             `bson:"LastName"`
	Email     string             `bson:"Email"`
	Password  string             `bson:"Password"`
	Phone     string             `bson:"Phone"`
	Teacher   *Teacher           `bson:"Teacher,omitempty"`
}

// Student is a struct that represents a student
type Student struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName"`
	LastName  string             `bson:"LastName"`
	Email     string             `bson:"Email"`
	Password  string             `bson:"Password"`
	Phone     string             `bson:"Phone"`
	Student   *Student           `bson:"Student,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
