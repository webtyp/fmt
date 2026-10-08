package main

import (
	"fmt"
	"strconv"
	"strings"
)

// processTextWithStandardLib simulates text processing using standard library
func processTextWithStandardLib(texts []string) []string {
	results := make([]string, len(texts))
	for i, text := range texts {
		replaced := strings.ToLower(text)

		// Capitalizar solo la primera letra del string completo
		out := replaced
		if len(out) > 0 {
			out = strings.ToUpper(out[:1]) + out[1:]
		}
		results[i] = out
	}
	return results
}

// processNumbersWithStandardLib simulates number processing
func processNumbersWithStandardLib(numbers []float64) []string {
	results := make([]string, len(numbers))
	for i, num := range numbers {
		// EQUIVALENT OPERATIONS: format with 2 decimals
		formatted := strconv.FormatFloat(num, 'f', 2, 64)
		results[i] = formatted
	}
	return results
}

func main() {
	fmt.Println("Standard library benchmark main function - used for testing only")
}
