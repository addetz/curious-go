package enums

/**
	You are working with integrating with a legacy external service for a major
	bookseller. The client sends you the following list of book categories  as part
	of the required JSON schema.
	----------------------------
	Novel : 1
	Fairy-tale : 2
	Drama : 3
	Poetry : 4
	Technical : 5
	Biography : 6
	----------------------------

	Read more: https://www.gopherguides.com/articles/how-to-use-iota-in-golang
**/

type BookCategory int
type BookFormat string

const (
	// New developer joins the team
	Print BookFormat = "Paperback"
	Ebook BookFormat = "Ebook"

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

func CreateBook(title string, category BookCategory, format BookFormat) *Book {
	return &Book{
		Title:    title,
		Category: category,
		Format:   format,
	}
}
