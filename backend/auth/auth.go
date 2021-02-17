package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-auth/utils"

	"golang.org/x/crypto/bcrypt"
)

// User represents the user entry in a fake database
type User struct {
	name string
	hash string
}

// Credentials represent auth credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]User, 10)

// Login takes typical Credentials and logs user in
func Login(w http.ResponseWriter, req *http.Request) {}

// Register takes typical Credentials and registers user in the (fake) database
func Register(w http.ResponseWriter, req *http.Request) {
	utils.SetupCorsResponse(&w, req)
	if req.Method == http.MethodOptions {
		return
	}

	var cr Credentials

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
	users[cr.Username] = User{cr.Username, string(hash)}

	// Log all the users
	for _, user := range users {
		fmt.Println(user)
	}

	w.WriteHeader(http.StatusCreated)
}
