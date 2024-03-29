package channels_test

import (
	"testing"

	enums "github.com/addetz/curious-go/channels"
	channels "github.com/addetz/curious-go/channels/ex1"
)

func TestAddStock(t *testing.T) {
	store := channels.NewStore()
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
			addRequest := channels.StockRequest{
				Book:  tc.book,
				Count: tc.increment,
				Type:  channels.AddStock,
			}
			stock := store.ReceiveRequest(addRequest)
			if stock.Count != tc.expectedCount {
				t.Fatalf("[%s]: incorrect returned count; got %d, want %d", name, stock.Count, tc.expectedCount)
			}
			getRequest := channels.StockRequest{
				Book: tc.book,
				Type: channels.GetStock,
			}
			retrievedStock := store.ReceiveRequest(getRequest)
			if retrievedStock.Count != tc.expectedCount {
				t.Fatalf("[%s]: incorrect retrieved count; got %d, want %d", name, retrievedStock.Count, tc.expectedCount)
			}
		})
	}
}
