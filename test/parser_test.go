package test

import (
	"fmt"
	. "go_jsonparser/api/lpjsonparser"
	"testing"
)

func TestParser(t *testing.T) {
	jsonStr := `{"name": "test", "age": 25}`

	value, err := Parse(jsonStr)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	fmt.Printf("Parsed value: %v\n", value)
}
