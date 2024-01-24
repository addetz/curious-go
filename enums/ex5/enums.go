package enums

type BookCategory int
type BookFormat string

const (
	Paperback BookFormat = "Paperback"
	Hardback  BookFormat = "Hardback"
	Kindle    BookFormat = "Kindle"
	PDF       BookFormat = "PDF"
)

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
	Format   BookFormat
}

func NewBook(title string, category BookCategory, format BookFormat) Book {
	return Book{
		Title:    title,
		Category: category,
		Format:   format,
	}
}
