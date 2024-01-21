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

func (s *Store) AddStock(newBook enums.Book, count int) Stock {
	stock, ok := s.stock[newBook.ID]
	if !ok {
		newStock := Stock{
			Book:  newBook,
			Count: count,
		}
		s.stock[newBook.ID] = newStock
		return newStock
	}

	newCount := stock.Count + count
	stock.Count = newCount
	return stock
}

func (s *Store) GetStock(id string) (*Stock, error) {
	stock, ok := s.stock[id]
	if !ok {
		return nil, fmt.Errorf("no book with id %s found", id)
	}

	return &stock, nil
}
