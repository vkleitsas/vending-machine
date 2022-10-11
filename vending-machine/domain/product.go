package domain

type Product struct {
	ID     int    `json:"id" pg:"id,pk"`
	Amount int    `json:"Amount" pg:"amount"`
	Cost   int    `json:"Cost" pg:"cost"`
	Name   string `json:"Name" pg:"name"`
	Seller int    `json:"Seller" pg:"seller"`
}
