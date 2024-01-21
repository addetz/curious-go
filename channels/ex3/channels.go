package channels

import (
	"sync"

	enums "github.com/addetz/curious-go/channels"
)

type StockRequestType int

type Stock struct {
	Book  enums.Book
	Count int
}

type StockRequest struct {
	Book  enums.Book
	Count int
}

type Store struct {
	stock    sync.Map
	requests chan StockRequest
	results  chan Stock
}

func NewStore() *Store {
	ns := &Store{
		requests: make(chan StockRequest),
		results:  make(chan Stock),
	}

	go ns.processRequests()
	go ns.processRequests()

	return ns
}

func (s *Store) ReceiveRequest(req StockRequest) Stock {
	s.requests <- req
	return <-s.results
}

func (s *Store) processRequests() {
	for req := range s.requests {
		stock := s.addStock(req.Book, req.Count)
		s.results <- stock
	}
}

func (s *Store) addStock(newBook enums.Book, count int) Stock {
	newStock := Stock{
		Book:  newBook,
		Count: count,
	}
	if stock, ok := s.stock.Load(newBook.ID); ok {
		existingStock := stock.(Stock)
		newStock.Count += existingStock.Count
	}

	s.stock.Store(newBook.ID, newStock)
	return newStock
}
