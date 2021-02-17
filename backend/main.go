package main

import (
	"log"
	"net/http"
	"simple-auth/auth"
)

func main() {
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/secret-data", secretData)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func secretData(res http.ResponseWriter, req *http.Request) {}
