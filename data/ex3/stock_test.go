package data_test

import (
	"testing"

	enums "github.com/addetz/curious-go/data"
	data "github.com/addetz/curious-go/data/ex3"
)

func TestAddStock(t *testing.T) {
	store := data.NewStore()
	testCases := map[string]struct {
		book          enums.Book
		increment     []int
		expectedCount int
	}{
		"empty store": {
			book:          enums.NewBook("The Go Programming Language", enums.Technical, enums.Paperback),
			increment:     []int{4},
			expectedCount: 4,
		},

		"new book": {
			book:          enums.NewBook("For Whom the Bell Tolls", enums.Novel, enums.Paperback),
			increment:     []int{3},
			expectedCount: 3,
		},

		"existing book": {
			book:          enums.NewBook("The Bell Jar", enums.Novel, enums.Paperback),
			increment:     []int{2, 4},
			expectedCount: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			for _, inc := range tc.increment {
				store.AddStock(tc.book, inc)
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
