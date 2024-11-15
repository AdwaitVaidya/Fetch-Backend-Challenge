package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	if !validateReceipt(&receipt) {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	receipt.ID = generateUniqueID()
	calculator := NewPointsCalculator(&receipt)
	receipt.PointsAwarded = calculator.CalculatePoints()
	receipts[receipt.ID] = &receipt

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receipt.ID})
}

func getReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, ok := receipts[id]
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"points": receipt.PointsAwarded})
}
