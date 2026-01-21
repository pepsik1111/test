package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request to create note")

	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	userReq, _ := json.Marshal(map[string]int{
		"id": int(data["user_id"].(float64)),
	})

	resp1, err := http.Post("http://user-service:8080/check", "application/json", bytes.NewBuffer(userReq))
	if err != nil {
		http.Error(w, "Error checking user", http.StatusInternalServerError)
		return
	}
	defer resp1.Body.Close()

	var userCheck map[string]bool
	json.NewDecoder(resp1.Body).Decode(&userCheck)

	if !userCheck["exists"] {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	bodyBytes, _ := json.Marshal(data)
	resp2, err := http.Post("http://note-service:8080/create", "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		http.Error(w, "Error creating note", http.StatusInternalServerError)
		return
	}
	defer resp2.Body.Close()

	var created map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&created)
	json.NewEncoder(w).Encode(created)
}

//yjdsq rvvbn

func main() {
	http.HandleFunc("/create-note", createNoteHandler)
	http.ListenAndServe(":8080", nil)
}
