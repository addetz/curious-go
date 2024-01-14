package enums_test

import (
	"log"
	"testing"

	enums "github.com/addetz/curious-go/tests/ex2"
)

func TestBookCategory(t *testing.T) {
	testCases := map[string]struct {
		title            string
		category         enums.BookCategory
		format           enums.BookFormat
		expectedCategory int
		expectedFormat   string
	}{
		"novel": {
			title:            "For Whom the Bell Tolls",
			category:         enums.Novel,
			format:           enums.Print,
			expectedCategory: 1,
			expectedFormat:   "Paperback",
		},

		"technical": {
			title:            "The Go Programming Language",
			category:         enums.Technical,
			format:           enums.Print,
			expectedCategory: 5,
			expectedFormat:   "Paperback",
		},

		"bio": {
			title:            "Becoming",
			category:         enums.Biography,
			format:           enums.Ebook,
			expectedCategory: 6,
			expectedFormat:   "Ebook",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			currentName, currentTc  := name, tc
			t.Parallel()
			book := enums.CreateBook(currentTc.title , currentTc.category, currentTc.format)
			log.Println(book)
			if book.Category != enums.BookCategory(currentTc.expectedCategory) {
				t.Fatalf("[%s]: incorrect category; got %d, want %d", currentName, book.Category, currentTc.expectedCategory)
			}
			if book.Format != enums.BookFormat(currentTc.expectedFormat) {
				t.Fatalf("[%s]: incorrect format; got %s, want %s", currentName, book.Format, currentTc.expectedFormat)
			}
		})
	}
}
