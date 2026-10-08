package fmt_test

import (
	"testing"
	. "webtyp.com/fmt"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple string",
			input:    "hello",
			expected: `"hello"`,
		},
		{
			name:     "String with spaces",
			input:    "hello world",
			expected: `"hello world"`,
		},
		{
			name:     "String with quotes",
			input:    `say "hello"`,
			expected: `"say \"hello\""`,
		},
		{
			name:     "String with backslash",
			input:    `path\to\file`,
			expected: `"path\\to\\file"`,
		},
		{
			name:     "String with newline",
			input:    "line1\nline2",
			expected: `"line1\nline2"`,
		},
		{
			name:     "String with tab",
			input:    "before\tafter",
			expected: `"before\tafter"`,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: `""`,
		},
		{
			name:     "String with carriage return",
			input:    "before\rafter",
			expected: `"before\rafter"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := Sprintf("%q", tt.input)
			if out != tt.expected {
				t.Errorf("Sprintf(%%q) = %s, want %s", out, tt.expected)
			}
		})
	}
}

// A non-string value is quoted in its string form (unlike Go, on purpose):
// what the old Convert(x).Quote() returned.
func TestQuoteNonString(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{123, `"123"`},
		{true, `"true"`},
		{1.5, `"1.5"`},
	}
	for _, c := range cases {
		if got := Sprintf("%q", c.in); got != c.want {
			t.Errorf("Sprintf(%%q, %v) = %s, want %s", c.in, got, c.want)
		}
	}
}
