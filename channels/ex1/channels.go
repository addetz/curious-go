package channels

import (
	"fmt"
	"log"

	enums "github.com/addetz/curious-go/channels"
)

type StockRequestType int

const (
	AddStock StockRequestType = iota
	GetStock
)

type Stock struct {
	Book  enums.Book
	Count int
}

type StockRequest struct {
	Book  enums.Book
	Count int
	Type  StockRequestType
}

type Store struct {
	stock    map[string]Stock
	requests chan StockRequest
	results  chan Stock
}

func NewStore() *Store {
	ns := &Store{
		stock:    make(map[string]Stock),
		requests: make(chan StockRequest),
		results:  make(chan Stock),
	}

	go ns.processRequests()

	return ns
}

func (s *Store) ReceiveRequest(req StockRequest) Stock {
	s.requests <- req
	return <-s.results
}

func (s *Store) processRequests() {
	for req := range s.requests {
		switch req.Type {
		case AddStock:
			stock := s.addStock(req.Book, req.Count)
			s.results <- stock
		case GetStock:
			stock, err := s.getStock(req.Book.ID)
			if err != nil {
				log.Println("processRequests:", err)
			}
			s.results <- *stock
		}
	}
}

func (s *Store) addStock(newBook enums.Book, count int) Stock {
	newStock := Stock{
		Book:  newBook,
		Count: count,
	}
	if stock, ok := s.stock[newBook.ID]; ok {
		newStock.Count += stock.Count
	}

	s.stock[newBook.ID] = newStock
	return newStock
}

func (s *Store) getStock(id string) (*Stock, error) {
	stock, ok := s.stock[id]
	if !ok {
		return nil, fmt.Errorf("no book with id %s found", id)
	}

	return &stock, nil
}
