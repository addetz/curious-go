package channels_test

import (
	"testing"

	channels "github.com/addetz/curious-go/channels/ex2"
	enums "github.com/addetz/curious-go/enums/ex6"
)

func TestAddStock(t *testing.T) {
	store := channels.NewStore()
	testCases := map[string]struct {
		book      enums.Book
		increment int
	}{
		"empty store": {
			book:      *enums.CreateBook("The Go Programming Language", enums.Technical, enums.Print),
			increment: 4,
		},

		"new book": {
			book:      *enums.CreateBook("For Whom the Bell Tolls", enums.Novel, enums.Print),
			increment: 3,
		},

		"existing book": {
			book:      *enums.CreateBook("The Go Programming Language", enums.Technical, enums.Print),
			increment: 2,
		},

		"existing book 2": {
			book:      *enums.CreateBook("The Go Programming Language", enums.Technical, enums.Print),
			increment: 10,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			currentName, currentTc := name, tc
			t.Parallel()
			stock := store.AddStock(currentTc.book, currentTc.increment)
			stockIncrement := stock.CurrentCount - stock.PreviousCount
			if stockIncrement != currentTc.increment {
				t.Fatalf("[%s]: incorrect increment; got %d, want %d", currentName, stockIncrement, currentTc.increment)
			}
		})
	}
}
