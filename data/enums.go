package enums

import "github.com/google/uuid"

type BookCategory int
type BookFormat string

const (
	Paperback BookFormat = "Paperback"
	Hardback  BookFormat = "Hardback"
	Kindle    BookFormat = "Kindle"
	PDF       BookFormat = "PDF"

	Novel     BookCategory = 1
	FairyTale BookCategory = 2
	Drama     BookCategory = 3
	Poetry    BookCategory = 4
	Technical BookCategory = 5
	Biography BookCategory = 6
)

type Book struct {
	ID       string
	Title    string
	Category BookCategory
	Format   BookFormat
}

func NewBook(title string, category BookCategory, format BookFormat) Book {
	id := uuid.New().String()
	return Book{
		ID:       id,
		Title:    title,
		Category: category,
		Format:   format,
	}
}
