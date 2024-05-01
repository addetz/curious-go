package enums

type BookCategory int

// Novel: 1
// Fairytale: 2
// Drama: 3
// Poetry: 4
// Technical: 5
// Biography: 6

const (
	Novel BookCategory = iota
	FairyTale
	Drama
	Poetry
	Technical
	Biography
)

type Book struct {
	Title    string
	Category BookCategory
}

func NewBook(title string, category BookCategory) Book {
	return Book{
		Title:    title,
		Category: category,
	}
}
