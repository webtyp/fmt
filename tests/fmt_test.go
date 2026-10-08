package fmt_test

import (
	"testing"
	. "webtyp.com/fmt"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{
			name:     "String formatting",
			format:   "Hello %s!",
			args:     []any{"World"},
			expected: "Hello World!",
		},
		{
			name:     "Integer formatting",
			format:   "Value: %d",
			args:     []any{42},
			expected: "Value: 42",
		},
		{
			name:     "Float formatting",
			format:   "Pi: %.2f",
			args:     []any{3.14159},
			expected: "Pi: 3.14",
		},
		{
			name:     "Multiple arguments",
			format:   "Hello %s, you have %d messages",
			args:     []any{"Alice", 5},
			expected: "Hello Alice, you have 5 messages",
		},
		{
			name:     "Binary formatting",
			format:   "Binary: %b",
			args:     []any{7},
			expected: "Binary: 111",
		},
		{
			name:     "Hexadecimal formatting",
			format:   "Hex: %x",
			args:     []any{255},
			expected: "Hex: ff",
		},
		{
			name:     "Octal formatting",
			format:   "Octal: %o",
			args:     []any{64},
			expected: "Octal: 100",
		},
		{
			name:     "Value formatting",
			format:   "Bool: %v",
			args:     []any{true},
			expected: "Bool: true",
		},
		{
			name:     "Percent sign",
			format:   "100%% complete",
			args:     []any{},
			expected: "100% complete",
		},
		{
			name:     "Missing argument",
			format:   "Value: %d",
			args:     []any{},
			expected: "",
		},
		{
			name:     "Boolean true formatting",
			format:   "Bool: %t",
			args:     []any{true},
			expected: "Bool: true",
		},
		{
			name:     "Quoted string formatting",
			format:   "Quoted: %q",
			args:     []any{"hello"},
			expected: "Quoted: \"hello\"",
		},
		{
			name:     "Pointer formatting",
			format:   "Pointer: %p",
			args:     []any{new(int)},
			expected: "Pointer: 0x",
		},
		{
			name:     "Unicode format",
			format:   "Unicode: %U",
			args:     []any{'A'},
			expected: "Unicode: U+0041",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := Sprintf(test.format, test.args...)
			if out != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, out)
			}
		})
	}
}

func TestErrorFormatting(t *testing.T) {
	err := Err("file not found").Error()
	want := "file not found"
	if err != want {
		t.Errorf("Err want %q got %q", want, err)
	}

	// Verify Err returns Conv that implements error
	var _ error = Err("some error")
}
