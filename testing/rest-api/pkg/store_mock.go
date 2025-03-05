package pkg

type MockStore struct {
	Items map[string]Item
}

func (m *MockStore) CreateItem(item Item) error {
	m.Items[item.ID] = item
	return nil
}
func (m *MockStore) GetAllItems() []Item {
	items := []Item{}
	for _, item := range m.Items {
		items = append(items, item)
	}
	return items
}
func (m *MockStore) GetItemByID(id string) (Item, bool) {
	item, found := m.Items[id]
	return item, found
}
func (m *MockStore) UpdateItem(id string, updated Item) bool {
	if _, exists := m.Items[id]; !exists {
		return false
	}
	m.Items[id] = updated
	return true
}
func (m *MockStore) DeleteItem(id string) bool {
	if _, exists := m.Items[id]; !exists {
		return false
	}
	delete(m.Items, id)
	return true
}
