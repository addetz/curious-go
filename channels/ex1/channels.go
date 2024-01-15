package channels

import (
	enums "github.com/addetz/curious-go/enums/ex6"
)

type Stock struct {
	CurrentCount  int
	PreviousCount int
	Book          enums.Book
}

func NewStock(book enums.Book, count int) *Stock {
	return &Stock{
		Book:         book,
		CurrentCount: count,
	}
}

type stockRequest struct {
	increment int
	book      enums.Book
}

type Store struct {
	stocks []*Stock
}

func NewStore() *Store {
	return &Store{
		stocks: make([]*Stock, 0),
	}
}

func (s *Store) AddStock(b enums.Book, c int) Stock {
	req := stockRequest{
		book:      b,
		increment: c,
	}

	for _, cs := range s.stocks {
		if cs.Book == req.book {
			newCount := cs.CurrentCount + req.increment
			cs.PreviousCount = cs.CurrentCount
			cs.CurrentCount = newCount
			return *cs
		}
	}
	newStock := NewStock(b, c)
	s.stocks = append(s.stocks, newStock)
	return *newStock
}
