package handlers

import (
	"encoding/json"
	"net/http"
)

type Attendance struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListAttendance(w http.ResponseWriter, _ *http.Request) {
	items := []Attendance{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
