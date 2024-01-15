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
		requests: make(chan *stockRequest, 1),
	}

	go ns.processRequests()
	go ns.processRequests()
	return ns
}

func (s *Store) AddStock(b enums.Book, c int) Stock {
	result := make(chan Stock)
	req := &stockRequest{
		book:      b,
		increment: c,
		result:    result,
	}
	s.requests <- req
	res := <-result
	return res
}

func (s *Store) processRequests() {
	for req := range s.requests {
		stock := s.handleUpdate(req)
		req.result <- stock
	}
}

func (s *Store) handleUpdate(req *stockRequest) Stock {
	for _, cs := range s.stocks {
		if cs.Book == req.book {
			newCount := cs.CurrentCount + req.increment
			cs.PreviousCount = cs.CurrentCount
			cs.CurrentCount = newCount
			return *cs
		}
	}
	newStock := NewStock(req.book, req.increment)
	s.stocks = append(s.stocks, newStock)
	return *newStock
}
