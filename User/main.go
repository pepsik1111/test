package main

import (
	"encoding/json"
	"net/http"
)

var users = map[int]string{
	1: "Alice",
	2: "Kostya",
	3: "NoName",
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID int `json:"id"`
	}
	type Response struct {
		Exists bool `json:"exssts"`
	}

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	_, ok := users[req.ID]

	json.NewEncoder(w).Encode(Response{Exists: ok})

}

func main() {
	http.HandleFunc("/check", checkUser)
	http.ListenAndServe(":8080", nil)
}
