package main

import (
	. "webtyp.com/fmt"
)

// processTextWithWebTyp simulates text processing using fmt (equivalent to standard lib)
func processTextWithWebTyp(texts []string) []string {
	results := make([]string, len(texts))
	for i, text := range texts {
		out := Convert(text).
			ToLower().
			Capitalize().
			String()
		results[i] = out
	}
	return results
}

// processNumbersWithWebTyp simulates number processing (equivalent to standard lib)
func processNumbersWithWebTyp(numbers []float64) []string {
	results := make([]string, len(numbers))
	for i, num := range numbers {
		// EQUIVALENT OPERATIONS: Same formatting as standard library
		formatted := Convert(num).
			Round(2).
			String()
		results[i] = formatted
	}
	return results
}

func main() {
	println("fmt benchmark main function - used for testing only")
}
