package cmd

import (
	"Advanced_programming_project/db"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func teachLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		phone := r.FormValue("phone")
		password := r.FormValue("password")

		collection := db.Client.Database("Learn").Collection("teachers")
		var teacher Teacher
		err := collection.FindOne(context.Background(), bson.M{"phone": phone}).Decode(&teacher)
		if err != nil {
			http.Error(w, "Invalid phone or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid phone or password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/teach/%s", teacher.ID.Hex()), http.StatusSeeOther)
	} else {
		renderTemplate(w, "vollogin.html", nil)
	}
}

func teachRegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		teacher := Teacher{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  string(hashedPassword),
			Phone:     phone,
		}

		collection := db.Client.Database("Learn").Collection("teachers")

		result, err := collection.InsertOne(context.Background(), teacher)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		insertedID := result.InsertedID.(primitive.ObjectID)
		http.Redirect(w, r, fmt.Sprintf("/teach/%s", insertedID.Hex()), http.StatusSeeOther)
	} else {
		renderTemplate(w, "volreg.html", nil)
	}
}

func studLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		phone := r.FormValue("phone")
		password := r.FormValue("password")

		collection := db.Client.Database("SantaWeb").Collection("children")
		var student Student
		err := collection.FindOne(context.Background(), bson.M{"phone": phone}).Decode(&student)
		if err != nil {
			errorResponse := ErrorResponse{Status: "400", Message: "Некорректное JSON-сообщение"}
			sendJSONResponse(w, errorResponse, http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
		if err != nil {
			errorResponse := ErrorResponse{Status: "400", Message: "Некорректное JSON-сообщение"}
			sendJSONResponse(w, errorResponse, http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/stud/%s", student.ID.Hex()), http.StatusSeeOther)
	} else {
		renderTemplate(w, "chilog.html", nil)
	}
}

func studRegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		student := Student{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Phone:     phone,
			Password:  string(hashedPassword),
		}

		collection := db.Client.Database("Learn").Collection("student")

		result, err := collection.InsertOne(context.Background(), student)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		insertedID := result.InsertedID.(primitive.ObjectID)
		http.Redirect(w, r, fmt.Sprintf("/stud/%s", insertedID.Hex()), http.StatusSeeOther)
	} else {
		renderTemplate(w, "chireg.html", nil)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("frontend/templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func volunteerPersonalPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacherID := vars["id"]

	var teacher Teacher
	collection := db.Client.Database("Learn").Collection("teachers")
	objID, _ := primitive.ObjectIDFromHex(teacherID)

	err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&teacher)
	if err != nil {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	renderTemplate(w, "vol.html", teacher)
}

func childPersonalPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studID := vars["id"]

	var student Student
	collection := db.Client.Database("Learn").Collection("students")
	objID, _ := primitive.ObjectIDFromHex(studID)

	err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&student)
	if err != nil {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	renderTemplate(w, "chil.html", studID)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
