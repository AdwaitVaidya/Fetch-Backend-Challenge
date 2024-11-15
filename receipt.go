package main

import (
	"fmt"
	"regexp"
	"time"
)

type Receipt struct {
	Retailer      string  `json:"retailer"`
	PurchaseDate  string  `json:"purchaseDate"`
	PurchaseTime  string  `json:"purchaseTime"`
	Items         []Item  `json:"items"`
	Total         float64 `json:"total,string"`
	ID           string  `json:"id"`
	PointsAwarded int64   `json:"points"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

var receipts = make(map[string]*Receipt)

func validateReceipt(receipt *Receipt) bool {
	return validateRetailer(receipt.Retailer) &&
		validatePurchaseDate(receipt.PurchaseDate) &&
		validatePurchaseTime(receipt.PurchaseTime) &&
		validateItems(receipt.Items) &&
		validateTotal(receipt.Total)
}

func validateRetailer(retailer string) bool {
	retailerRegex := regexp.MustCompile(`^[\w\s\-&]+$`)
	return retailerRegex.MatchString(retailer)
}

func validatePurchaseDate(purchaseDate string) bool {
	_, err := time.Parse("2006-01-02", purchaseDate)
	return err == nil
}

func validatePurchaseTime(purchaseTime string) bool {
	_, err := time.Parse("15:04", purchaseTime)
	return err == nil
}

func validateItems(items []Item) bool {
	if len(items) < 1 {
		return false
	}
	for _, item := range items {
		itemDescriptionRegex := regexp.MustCompile(`^[\w\s\-]+$`)
		priceRegex := regexp.MustCompile(`^\d+\.\d{2}$`)
		if !itemDescriptionRegex.MatchString(item.ShortDescription) || !priceRegex.MatchString(fmt.Sprintf("%.2f", item.Price)) {
			return false
		}
	}
	return true
}

func validateTotal(total float64) bool {
	totalRegex := regexp.MustCompile(`^\d+\.\d{2}$`)
	return totalRegex.MatchString(fmt.Sprintf("%.2f", total))
}

func generateUniqueID() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())
}
