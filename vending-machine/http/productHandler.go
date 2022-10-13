package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"vending-machine/app"
	"vending-machine/domain"
)

type ProductHandler struct {
	service app.ProductCrudService
}

func NewProductHandler(s app.ProductCrudService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestUser := r.Context().Value("requestUser").(domain.User)
	log.Println("p.seller " + strconv.Itoa(request.Seller) + "request user " + strconv.Itoa(requestUser.ID))
	result, err := h.service.CreateProduct(request, requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *ProductHandler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestUser := r.Context().Value("requestUser").(domain.User)
	err = h.service.DeleteProduct(request, requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) EditProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestUser := r.Context().Value("requestUser").(domain.User)
	result, err := h.service.UpdateProduct(request, requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
