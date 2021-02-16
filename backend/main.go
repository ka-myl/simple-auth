package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// User represents the user entry in a fake database
type user struct {
	name string
	hash string
}

// Credentials represent auth credentials
type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]user, 10)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/secret-data", secretData)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func login(w http.ResponseWriter, req *http.Request) {}

func register(w http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&w, req)
	if req.Method == http.MethodOptions {
		return
	}

	var cr credentials

	// Decode json payload from request
	err := json.NewDecoder(req.Body).Decode(&cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate hashed password
	hash, err := bcrypt.GenerateFromPassword([]byte(cr.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save user to fake database
	users[cr.Username] = user{cr.Username, string(hash)}

	// Log all the users
	for _, user := range users {
		fmt.Println(user)
	}

	w.WriteHeader(http.StatusCreated)
}

func secretData(res http.ResponseWriter, req *http.Request) {}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
