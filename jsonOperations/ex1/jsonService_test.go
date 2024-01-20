package jsonoperations_test

import (
	"fmt"
	"testing"

	enums "github.com/addetz/curious-go/enums/ex6"
	jsonoperations "github.com/addetz/curious-go/jsonOperations/ex1"
)

func TestMarshallBook(t *testing.T) {
	testCases := map[string]struct {
		title    string
		category enums.BookCategory
		format   enums.BookFormat
	}{
		"novel": {
			title:    "For Whom the Bell Tolls",
			category: enums.Novel,
			format:   enums.Print,
		},

		"technical": {
			title:    "The Go Programming Language",
			category: enums.Technical,
			format:   enums.Print,
		},

		"bio": {
			title:    "Becoming",
			category: enums.Biography,
			format:   enums.Ebook,
		},

		"novel ampersand": {
			title:    "The Daughter of Smoke & Bone",
			category: enums.Novel,
			format:   enums.Ebook,
		},
	}

	getExpectedMarshalledBook := func(b *enums.Book) string {
		return fmt.Sprintf("{\"title\":\"%s\",\"category\":%d,\"format\":\"%s\"}",
			b.Title, b.Category, b.Format)
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			book := enums.CreateBook(tc.title, tc.category, tc.format)
			marshalled, err := jsonoperations.MarshallBook(book)
			if err != nil {
				t.Fatalf("[%s]: unexpected error: %v", name, err)
			}
			output, expectedOutput := string(marshalled), getExpectedMarshalledBook(book)
			if output != expectedOutput {
				t.Fatalf("[%s]: incorrect marshalled output; got %s, want %s", name, output, expectedOutput)
			}
		})
	}
}
