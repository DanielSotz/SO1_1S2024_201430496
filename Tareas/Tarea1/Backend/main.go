// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func userDataHandler(w http.ResponseWriter, r *http.Request) {
	// Mock data for demonstration purposes
	user := UserData{
		Name:  "John Doe",
		Age:   25,
		Email: "john.doe@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/data", userDataHandler)
	http.ListenAndServe(":8081", nil)
}
