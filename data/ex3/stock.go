package data

import (
	"fmt"

	enums "github.com/addetz/curious-go/data"
)

type Stock struct {
	Book  enums.Book
	Count int
}

type Store struct {
	stock map[string]Stock
}

func NewStore() *Store {
	return &Store{
		stock: make(map[string]Stock),
	}
}

func (s *Store) AddStock(newBook enums.Book, count int) {
	existingStock := 0
	if stock, ok := s.stock[newBook.ID]; ok {
		existingStock = stock.Count
	}

	s.stock[newBook.ID] = Stock{
		Book:  newBook,
		Count: count + existingStock,
	}
}

func (s *Store) GetStock(id string) (*Stock, error) {
	stock, ok := s.stock[id]
	if !ok {
		return nil, fmt.Errorf("no book with id %s found", id)
	}

	return &stock, nil
}
