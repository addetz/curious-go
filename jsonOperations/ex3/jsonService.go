package jsonoperations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	enums "github.com/addetz/curious-go/enums/ex6"
)

func MarshallBook(b *enums.Book) ([]byte, error) {
	marshalledBook := new(bytes.Buffer)
	enc := json.NewEncoder(marshalledBook)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(b); err != nil {
		return nil, fmt.Errorf("MarshallBook %v: %v", b, err)
	}

	out := strings.Trim(marshalledBook.String(), "\n")

	return []byte(out), nil
}
