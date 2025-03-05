package pkg

import "sync"

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ItemStore interface {
	CreateItem(item Item) error
	GetAllItems() []Item
	GetItemByID(id string) (Item, bool)
	UpdateItem(id string, updated Item) bool
	DeleteItem(id string) bool
}

type InMemoryStore struct {
	mu    sync.Mutex
	items map[string]Item
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		items: make(map[string]Item),
	}
}

func (s *InMemoryStore) CreateItem(item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.items[item.ID]; exists {
		return nil
	}
	s.items[item.ID] = item
	return nil
}

func (s *InMemoryStore) GetAllItems() []Item {
	s.mu.Lock()
	defer s.mu.Unlock()
	items := []Item{}
	for _, item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *InMemoryStore) GetItemByID(id string) (Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, found := s.items[id]
	return item, found
}

func (s *InMemoryStore) UpdateItem(id string, updated Item) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.items[id]; !exists {
		return false
	}
	s.items[id] = updated
	return true
}

func (s *InMemoryStore) DeleteItem(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.items[id]; !exists {
		return false
	}
	delete(s.items, id)
	return true
}
