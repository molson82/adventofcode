package models

type Set struct {
	items map[int]struct{}
}

// NewSet creates a new Set
func NewSet() *Set {
	return &Set{
		items: make(map[int]struct{}),
	}
}

// Add adds a new item to the Set. If the item already exists, it does nothing.
func (s *Set) Add(item int) {
	s.items[item] = struct{}{}
}

// Find checks if an item is in the Set.
func (s *Set) Find(item int) bool {
	_, exists := s.items[item]
	return exists
}

// Remove removes an item from the Set. If the item does not exist, it does nothing.
func (s *Set) Remove(item int) {
	delete(s.items, item)
}

// Clear removes all items from the Set.
func (s *Set) Clear() {
	s.items = make(map[int]struct{})
}

// Size returns the number of items in the Set.
func (s *Set) Size() int {
	return len(s.items)
}
