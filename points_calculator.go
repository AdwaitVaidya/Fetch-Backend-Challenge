package main

import (
	"math"
	"strings"
	"time"
	"unicode"
)

type PointsCalculator struct {
	receipt *Receipt
}

func NewPointsCalculator(receipt *Receipt) *PointsCalculator {
	return &PointsCalculator{
		receipt: receipt,
	}
}

func (pc *PointsCalculator) CalculatePoints() int64 {
	var totalPoints int64

	totalPoints += pc.calculateRetailerNamePoints()
	totalPoints += pc.calculateRoundDollarPoints()
	totalPoints += pc.calculateMultipleOf25Points()
	totalPoints += pc.calculateItemPairPoints()
	totalPoints += pc.calculateItemDescriptionPoints()
	totalPoints += pc.calculatePurchaseDatePoints()
	totalPoints += pc.calculatePurchaseTimePoints()

	return totalPoints
}

func (pc *PointsCalculator) calculateRetailerNamePoints() int64 {

    count:= 0
    for _, char := range pc.receipt.Retailer {
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            count++
        }
    }
    return int64(count)
}

func (pc *PointsCalculator) calculateRoundDollarPoints() int64 {
	if pc.receipt.Total == float64(int64(pc.receipt.Total)) {
		return 50
	}
	return 0
}

func (pc *PointsCalculator) calculateMultipleOf25Points() int64 {
	// Convert to cents to avoid floating point precision issues
	cents := int64(pc.receipt.Total * 100)
	if cents%25 == 0 {
		return 25
	}
	return 0
}

func (pc *PointsCalculator) calculateItemPairPoints() int64 {
	pairs := len(pc.receipt.Items) / 2
	return int64(pairs * 5)
}

func (pc *PointsCalculator) calculateItemDescriptionPoints() int64 {
	var points int64
	for _, item := range pc.receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			points += int64(math.Ceil(item.Price * 0.2))
		}
	}
	return points
}

func (pc *PointsCalculator) calculatePurchaseDatePoints() int64 {
	purchaseDate, _ := time.Parse("2006-01-02", pc.receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		return 6
	}
	return 0
}

func (pc *PointsCalculator) calculatePurchaseTimePoints() int64 {
	purchaseTime, _ := time.Parse("15:04", pc.receipt.PurchaseTime)
	targetStart, _ := time.Parse("15:04", "14:00")
	targetEnd, _ := time.Parse("15:04", "16:00")

	purchaseTimeToday := time.Date(0, 1, 1, purchaseTime.Hour(), purchaseTime.Minute(), 0, 0, time.UTC)
	if purchaseTimeToday.After(targetStart) && purchaseTimeToday.Before(targetEnd) {
		return 10
	}
	return 0
}
