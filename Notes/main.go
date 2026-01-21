package main

import (
	"encoding/json"
	"net/http"
)

type Note struct {
	ID     int    `json:"id"`
	UserID int    `json:"userid"`
	Text   string `json:"text"`
}

var (
	notes  []Note
	NextId int = 1
)

func createNote(w http.ResponseWriter, r *http.Request) {
	var n Note
	json.NewDecoder(r.Body).Decode(&n)
	notes = append(notes, n) 
	n.ID = NextId
	json.NewEncoder(w).Encode(n)
}	

func getNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func main() {
	http.HandleFunc("/create", createNote)
	http.HandleFunc("/list", getNotes)
	http.ListenAndServe(":8080", nil)
}
