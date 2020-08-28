package api

import (
	"encoding/json"
	"fmt"
)

type JSONError struct {
	Msg string
}

func (err JSONError) Error() string {
	val, _ := json.Marshal(err.Msg)
	return fmt.Sprintf("{\"detail\": %s } ", string(val))
}

func JSONResult(res string) string {
	val, _ := json.Marshal(res)
	return fmt.Sprintf("{\"result\": %s } ", string(val))
}

func FirstWords(value string, count int) string {
	// Loop over all indexes in the string.
	for i := range value {
		// If we encounter a space, reduce the count.
		if value[i] == ' ' {
			count -= 1
			// When no more words required, return a substring.
			if count == 0 {
				return value[0:i]
			}
		}
	}
	// Return the entire string.
	return value
}
