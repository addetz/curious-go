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
	stocks   []*Stock
	requests chan *stockRequest
	results  chan Stock
}

func NewStore() *Store {
	ns := &Store{
		stocks:   make([]*Stock, 0),
		requests: make(chan *stockRequest),
		results:  make(chan Stock),
	}

	go ns.processRequests()
	return ns
}

func (s *Store) AddStock(b enums.Book, c int) Stock {
	req := &stockRequest{
		book:      b,
		increment: c,
	}
	s.requests <- req
	res := <-s.results
	return res
}

func (s *Store) processRequests() {
	for req := range s.requests {
		for _, cs := range s.stocks {
			if cs.Book == req.book {
				newCount := cs.CurrentCount + req.increment
				cs.PreviousCount = cs.CurrentCount
				cs.CurrentCount = newCount
				s.results <- *cs
				return
			}
		}
		newStock := NewStock(req.book, req.increment)
		s.stocks = append(s.stocks, newStock)
		s.results <- *newStock
	}
}
