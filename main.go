package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getReceiptPoints).Methods("GET")

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", r)
}
