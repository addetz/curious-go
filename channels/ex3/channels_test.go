package channels_test

import (
	"testing"

	enums "github.com/addetz/curious-go/channels"
	channels "github.com/addetz/curious-go/channels/ex3"
)

func TestAddStock(t *testing.T) {
	store := channels.NewStore()
	testBook := enums.CreateBook("The Go Programming Language", enums.Technical, enums.Paperback)
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
			book:          enums.CreateBook("For Whom the Bell Tolls", enums.Novel, enums.Paperback),
			increment:     3,
			expectedCount: 3,
		},

		"existing book": {
			book:          testBook,
			increment:     2,
			expectedCount: 6,
		},

		"existing book 2 ": {
			book:          testBook,
			increment:     10,
			expectedCount: 16,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			addRequest := channels.StockRequest{
				Book:  tc.book,
				Count: tc.increment,
			}
			stock := store.ReceiveRequest(addRequest)
			if stock.Count != tc.expectedCount {
				t.Fatalf("[%s]: incorrect returned count; got %d, want %d", name, stock.Count, tc.expectedCount)
			}
		})
	}
}
