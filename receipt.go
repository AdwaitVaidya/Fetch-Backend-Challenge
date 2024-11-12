package main


type Receipt struct{

    Retailer string //'json:"retailer"'
    PurchaseDate string // 'json:"purchaseDate"'
    Item []Item //'json:"items"'
    Total float64 // 'json:"total"'
    ID string
    PointsAwarded int64
}

type Item struct {


    ShortDescription string
    Price float64
}



