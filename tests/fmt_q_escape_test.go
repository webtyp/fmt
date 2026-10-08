package fmt_test

import (
	"testing"
	. "webtyp.com/fmt"
)

// TestSprintfQ_EscapesSpecialChars: %q escapes quotes, backslashes and the
// common control characters (regression: it once wrapped the string unescaped).
func TestSprintfQ_EscapesSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain string",
			input:    "hello",
			expected: `"hello"`,
		},
		{
			name:     "double quotes must be escaped",
			input:    `{"key":"value"}`,
			expected: `"{\"key\":\"value\"}"`,
		},
		{
			name:     "backslash must be escaped",
			input:    `path\file`,
			expected: `"path\\file"`,
		},
		{
			name:     "newline must be escaped",
			input:    "line1\nline2",
			expected: `"line1\nline2"`,
		},
		{
			name:     "tab must be escaped",
			input:    "a\tb",
			expected: `"a\tb"`,
		},
		{
			// Caso real que causó el footer vacío en la TUI:
			// daemon.go usaba fmt.Sprintf(`...,"result":%q`, stateJSON)
			// donde stateJSON es un array JSON con comillas internas.
			name:     "json array payload (real MCP state bug)",
			input:    `[{"tab_title":"BUILD","handler_name":"WasmClient","handler_type":1}]`,
			expected: `"[{\"tab_title\":\"BUILD\",\"handler_name\":\"WasmClient\",\"handler_type\":1}]"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sprintf("%q", tt.input)
			if got != tt.expected {
				t.Errorf("Sprintf(%%q, %q)\n got:  %s\n want: %s", tt.input, got, tt.expected)
			}
		})
	}
}

// TestSprintfQ_ControlBytes: %q must produce a valid Go literal. Control bytes
// used to pass through raw, so an invisible byte reached logs.
func TestSprintfQ_ControlBytes(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"a\x01b", `"a\x01b"`},
		{"bell\a", `"bell\a"`},
		{"\b\f\v", `"\b\f\v"`},
		{"x\x7f", `"x\x7f"`},
		{"esc\x1b", `"esc\x1b"`},
		{"tab\there", `"tab\there"`},
		{"ñandú", `"ñandú"`}, // printable UTF-8 stays as is, as in Go
	}
	for _, tt := range tests {
		if got := Sprintf("%q", tt.input); got != tt.expected {
			t.Errorf("Sprintf(%%q, %q) = %s, want %s", tt.input, got, tt.expected)
		}
	}
}
