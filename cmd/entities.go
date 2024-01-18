package cmd

import "go.mongodb.org/mongo-driver/bson/primitive"

// Teacher is a struct that represents a teacher
type Teacher struct {
	ID        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Teacher   *Teacher
}

// Student is a struct that represents a student
type Student struct {
	ID        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Student   *Student
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
