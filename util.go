package main

import (
	"encoding/json"
	"fmt"
)

func JSONError(msg string) string {
	val, _ := json.Marshal(msg)
	return fmt.Sprintf("{\"detail\": %s } ", string(val))
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
