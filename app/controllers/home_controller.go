package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edrywn/toko-online/app/item"
	"github.com/gorilla/mux"
)

type itemHandler struct {
	itemService item.Service
}

func NewItemHandler(itemService item.Service) *itemHandler {
	return &itemHandler{itemService}

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func (h *itemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, _ := h.itemService.FindAll()

	var itemsResponse []item.ItemResponse
	for _, i := range items {
		itemResponse := convertToItemResponse(i)

		itemsResponse = append(itemsResponse, itemResponse)
	}

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(itemsResponse)
}
func (h *itemHandler) GetItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["ID"])

	i, err := h.itemService.FindByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemResponse := convertToItemResponse(i)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(itemResponse)

}

func (h *itemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")

	var itemRequest item.ItemRequest

	json.NewDecoder(r.Body).Decode(&itemRequest)

	item, err := h.itemService.Create(itemRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(item)

}

func (h *itemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")

	var itemRequest item.ItemRequest
	json.NewDecoder(r.Body).Decode(&itemRequest)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["ID"])
	i, err := h.itemService.Update(id, itemRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemResponse := convertToItemResponse(i)
	json.NewEncoder(w).Encode(itemResponse)
}
func (h *itemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["ID"])
	item, err := h.itemService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(item)

}

func convertToItemResponse(i item.Item) item.ItemResponse {
	return item.ItemResponse{
		ID:    i.ID,
		Name:  i.Name,
		Price: i.Price,
		Stock: i.Stock,
	}
}
