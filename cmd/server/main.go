package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/oak-ledger-app/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/attendance", handlers.ListAttendance)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
