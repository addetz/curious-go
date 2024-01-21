package jsonoperations

import (
	"encoding/json"
	"fmt"

	enums "github.com/addetz/curious-go/json_operations"
)

func MarshallBook(b enums.Book) ([]byte, error) {
	marshalledBook, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("MarshallBook %v: %v", b, err)
	}
	return marshalledBook, nil
}
