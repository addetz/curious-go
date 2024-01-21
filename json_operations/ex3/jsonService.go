package jsonoperations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	enums "github.com/addetz/curious-go/json_operations"
)

func MarshallBook(b *enums.Book) ([]byte, error) {
	// Create our own custom encoder
	marshalledBook := new(bytes.Buffer)
	enc := json.NewEncoder(marshalledBook)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(b); err != nil {
		return nil, fmt.Errorf("MarshallBook %v: %v", b, err)
	}

	// Trim the new lines
	out := strings.Trim(marshalledBook.String(), "\n")

	return []byte(out), nil
}
