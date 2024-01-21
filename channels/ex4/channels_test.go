package channels_test

import (
	"testing"

	channels "github.com/addetz/curious-go/channels/ex4"
	enums "github.com/addetz/curious-go/enums/ex6"
)

const (
	RUN_COUNT int = 10
)

func TestAddStock(t *testing.T) {
	store := channels.NewStore()
	testCases := map[string]struct {
		book      enums.Book
		increment int
	}{
		"empty store": {
			book:      enums.CreateBook("The Go Programming Language", enums.Technical, enums.Paperback),
			increment: 4,
		},

		"new book": {
			book:      enums.CreateBook("For Whom the Bell Tolls", enums.Novel, enums.Paperback),
			increment: 3,
		},

		"existing book": {
			book:      enums.CreateBook("The Go Programming Language", enums.Technical, enums.Paperback),
			increment: 2,
		},

		"existing book 2": {
			book:      enums.CreateBook("The Go Programming Language", enums.Technical, enums.Paperback),
			increment: 10,
		},

		"existing book 3": {
			book:      enums.CreateBook("The Go Programming Language", enums.Technical, enums.Paperback),
			increment: 1,
		},
	}
	// increase chances of error
	for i := 0; i < RUN_COUNT; i++ {
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
}
