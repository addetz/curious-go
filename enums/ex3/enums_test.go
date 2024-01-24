package enums_test

import (
	"testing"

	enums "github.com/addetz/curious-go/enums/ex3"
)

func TestBookCategory(t *testing.T) {
	testCases := map[string]struct {
		title            string
		category         enums.BookCategory
		expectedCategory int
	}{
		"novel": {
			title:            "For Whom the Bell Tolls",
			category:         enums.Novel,
			expectedCategory: 1,
		},

		"technical": {
			title:            "The Go Programming Language",
			category:         enums.Technical,
			expectedCategory: 5,
		},

		"bio": {
			title:            "Becoming",
			category:         enums.Biography,
			expectedCategory: 6,
		},

		"novel special character": {
			title:            "The Daughter of Smoke & Bone",
			category:         enums.Novel,
			expectedCategory: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			book := enums.NewBook(tc.title, tc.category)
			if book.Category != enums.BookCategory(tc.expectedCategory) {
				t.Fatalf("[%s]: incorrect category; got %d, want %d", name, book.Category, tc.expectedCategory)
			}
		})
	}
}
