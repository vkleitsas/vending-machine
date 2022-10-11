package domain

type PurchaseRequestDTO struct {
	ProductID int `json:"ProductID"`
	Amount    int `json:"Amount"`
	BuyerID   int `json:"BuyerID"`
}

type PurchaseResponseDTO struct {
	ProductID   int         `json:"ProductID"`
	ProductName string      `json:"Name"`
	BuyerID     int         `json:"BuyerID"`
	Change      map[int]int `json:"Change"`
}
