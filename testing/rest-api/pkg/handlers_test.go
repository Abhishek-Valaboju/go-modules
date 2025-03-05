package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateItem(t *testing.T) {
	store := &MockStore{Items: make(map[string]Item)}
	handler := NewHandler(store)
	item := Item{ID: "1", Name: "Laptop", Price: 999.99}
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	handler.CreateItem(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}
}
func TestGetAllItems(t *testing.T) {
	store := &MockStore{Items: map[string]Item{"1": {ID: "1", Name: "Laptop", Price: 999.99}}}
	handler := NewHandler(store)
	req, _ := http.NewRequest("GET", "/items", nil)
	rec := httptest.NewRecorder()
	handler.GetAllItems(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}
