package enums

type BookCategory int

const (
	Novel BookCategory = iota + 1
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

func CreateBook(title string, category BookCategory) Book {
	return Book{
		Title:    title,
		Category: category,
	}
}
