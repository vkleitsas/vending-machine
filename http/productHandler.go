package http

import (
	"encoding/json"
	"net/http"
	"vending-machine/app"
	"vending-machine/domain"
)

type CreateProductHandler struct {
	service app.ProductCrudService
}

func NewCreateProductHandler(s app.ProductCrudService) *CreateProductHandler {
	return &CreateProductHandler{
		service: s,
	}
}

func (h *CreateProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
