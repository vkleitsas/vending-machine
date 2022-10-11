package http

import (
	"encoding/json"
	"net/http"
	"vending-machine/app"
	"vending-machine/domain"
)

type TransactionsHandler struct {
	service app.TransactionsCrudService
}

func NewTransactionsHandler(s app.TransactionsCrudService) *TransactionsHandler {
	return &TransactionsHandler{
		service: s,
	}
}

func (h *TransactionsHandler) Purchase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.PurchaseRequestDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.service.Purchase(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
