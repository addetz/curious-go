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
	result    chan Stock
}

type Store struct {
	stocks   []*Stock
	requests chan *stockRequest
}

func NewStore() *Store {
	ns := &Store{
		stocks:   make([]*Stock, 0),
		requests: make(chan *stockRequest),
	}

	go ns.processRequests()
	return ns
}

func (s *Store) AddStock(b enums.Book, c int) Stock {
	result := make(chan Stock)
	sr := &stockRequest{
		book:      b,
		increment: c,
		result:    result,
	}
	s.requests <- sr
	res := <-result
	return res	
}

func (s *Store) processRequests() {
	for req := range s.requests {
		for _, cs := range s.stocks {
			if cs.Book == req.book {
				newCount := cs.CurrentCount + req.increment
				cs.PreviousCount = cs.CurrentCount
				cs.CurrentCount = newCount
				req.result <- *cs
				return
			}
		}
		newStock := NewStock(req.book, req.increment)
		s.stocks = append(s.stocks, newStock)
		req.result <- *newStock
	}
}
