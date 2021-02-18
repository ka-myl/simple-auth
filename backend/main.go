package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-auth/auth"
	"simple-auth/utils"
)

type secretData struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/secret", secret)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func secret(w http.ResponseWriter, req *http.Request) {
	utils.SetupCorsResponse(&w, req)
	if req.Method != http.MethodGet {
		return
	}

	if !auth.IsAuthenticated(req) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bs, err := json.Marshal(secretData{"Hello, this is secret message!"})
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}
