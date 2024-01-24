package data_test

import (
	"testing"

	enums "github.com/addetz/curious-go/data"
	data "github.com/addetz/curious-go/data/ex2"
)

func TestAddStock(t *testing.T) {
	store := data.NewStore()
	testBook := enums.NewBook("The Go Programming Language", enums.Technical, enums.Paperback)
	testCases := map[string]struct {
		book          enums.Book
		increment     int
		expectedCount int
	}{
		"empty store": {
			book:          testBook,
			increment:     4,
			expectedCount: 4,
		},

		"new book": {
			book:          enums.NewBook("For Whom the Bell Tolls", enums.Novel, enums.Paperback),
			increment:     3,
			expectedCount: 3,
		},

		"existing book": {
			book:          testBook,
			increment:     2,
			expectedCount: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stock := store.AddStock(tc.book, tc.increment)
			if stock.Count != tc.expectedCount {
				t.Fatalf("[%s]: incorrect returned count; got %d, want %d", name, stock.Count, tc.expectedCount)
			}
			retrievedStock, err := store.GetStock(tc.book.ID)
			if err != nil {
				t.Fatalf("[%s]: unexpected error; %v", name, err)
			}
			if retrievedStock.Count != tc.expectedCount {
				t.Fatalf("[%s]: incorrect retrieved count; got %d, want %d", name, retrievedStock.Count, tc.expectedCount)
			}
		})
	}
}
