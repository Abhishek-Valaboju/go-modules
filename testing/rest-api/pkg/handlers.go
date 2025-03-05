package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store ItemStore
}

func NewHandler(store ItemStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	h.store.CreateItem(item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
func (h *Handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.store.GetAllItems())
}
func (h *Handler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, found := h.store.GetItemByID(id)
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated Item
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if !h.store.UpdateItem(id, updated) {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}
func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if !h.store.DeleteItem(id) {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
