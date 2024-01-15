package enums_test

import (
	"testing"

	enums "github.com/addetz/curious-go/enums/ex5"
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
			book := enums.CreateBook(tc.title, tc.category, tc.format)
			if book.Category != enums.BookCategory(tc.expectedCategory) {
				t.Fatalf("[%s]: incorrect category; got %d, want %d", name, book.Category, tc.expectedCategory)
			}
			if book.Format != enums.BookFormat(tc.expectedFormat) {
				t.Fatalf("[%s]: incorrect format; got %s, want %s", name, book.Format, tc.expectedFormat)
			}
		})
	}
}
