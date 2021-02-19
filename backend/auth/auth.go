package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-auth/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// Claims represents JWT payload
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("some_random_thing")
var users = make(map[string]User, 10)

// Login takes typical Credentials and logs user in
func Login(w http.ResponseWriter, req *http.Request) {
	utils.SetupCorsResponse(&w, req)
	if req.Method != http.MethodPost {
		return
	}

	var cr Credentials

	// Decode json payload from request
	err := json.NewDecoder(req.Body).Decode(&cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Check if given user exists
	user, ok := users[cr.Username]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.hash), []byte(cr.Password))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create JWT
	exp := time.Now().Add(5 * time.Minute)
	c := Claims{
		Username: cr.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ts, err := t.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    ts,
		MaxAge:   60 * 5,
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
}

// Register takes typical Credentials and registers user in the (fake) database
func Register(w http.ResponseWriter, req *http.Request) {
	utils.SetupCorsResponse(&w, req)
	if req.Method != http.MethodPost {
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

// IsAuthenticated checks if a request comes from from authenticated user
func IsAuthenticated(req *http.Request) bool {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := req.Cookie("auth")
	if err != nil {
		return false
	}

	ts := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(ts, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return false
	}

	return true
}
