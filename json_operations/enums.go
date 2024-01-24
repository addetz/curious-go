package enums

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
	Title    string       `json:"title"`
	Category BookCategory `json:"category"`
	Format   BookFormat   `json:"format"`
}

func NewBook(title string, category BookCategory, format BookFormat) Book {
	return Book{
		Title:    title,
		Category: category,
		Format:   format,
	}
}
