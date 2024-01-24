package enums

type BookCategory int

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
