package main

import (

    "encoding/json"
    "net/http"
)

func processReceipt(w http.ResponseWriter, r *http.Request){

    var receipt Receipt
    err:= json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil{

        http.Error(w,"Receipt Decode Error", http.StatusBadRequest)
        return
    }

    if !validateReceipt(&receipt){

        http.Error(w,"Invalid Receipt Data",http.StatusBadRequest)
    }

    receipt.ID = generateUniqueID()
    receipt[receipt.ID] = &receipt

    receipt.PointsAwarded = calculatePoints(&receipt)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"id":receipt.ID})
}
