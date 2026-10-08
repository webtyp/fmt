package fmt_test

import (
	"testing"
	. "webtyp.com/fmt"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		maxWidth      any
		reservedChars any
		useReserved   bool
		expected      string
	}{
		{
			name:        "Conv shorter than max width",
			input:       "Hello",
			maxWidth:    10,
			useReserved: false,
			expected:    "Hello", // No padding expected
		},
		{
			name:        "Conv longer than max width with ellipsis",
			input:       "Hello, World!",
			maxWidth:    10,
			useReserved: false,
			expected:    "Hello, ...",
		},
		{
			name:          "Conv longer with reserved chars",
			input:         "Hello, World!",
			maxWidth:      10,
			reservedChars: 3,
			useReserved:   true,
			expected:      "Hell...",
		},
		{
			name:        "Very short max width",
			input:       "Hello",
			maxWidth:    2,
			useReserved: false,
			expected:    "He",
		},
		{
			name:        "maxWith is zero",
			input:       "Test",
			maxWidth:    0,
			useReserved: false,
			expected:    "Test",
		},
		{
			name:        "Empty string",
			input:       "",
			maxWidth:    5,
			useReserved: false,
			expected:    "", // No padding expected
		},
		{
			name:        "With uint8 maxWidth",
			input:       "Hello",
			maxWidth:    uint8(8),
			useReserved: false,
			expected:    "Hello", // No padding expected
		},
		{
			name:          "With uint16 maxWidth and uint8 reservedChars",
			input:         "Hello, World!",
			maxWidth:      uint16(10),
			reservedChars: uint8(3),
			useReserved:   true,
			expected:      "Hell...",
		},
		{
			name:          "With float32 maxWidth and int64 reservedChars",
			input:         "Hello, World!",
			maxWidth:      float32(10.5), // Should truncate to 10
			reservedChars: int64(2),
			useReserved:   true,
			expected:      "Hello...",
		},
		{
			name:        "With float64 maxWidth",
			input:       "Testing",
			maxWidth:    float64(9.9), // Should truncate to 9
			useReserved: false,
			expected:    "Testing", // No padding expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out string
			if tt.useReserved {
				out = Convert(tt.input).Truncate(tt.maxWidth, tt.reservedChars).String()
			} else {
				out = Convert(tt.input).Truncate(tt.maxWidth).String()
			}

			if out != tt.expected {
				t.Errorf("Convert(%q).Truncate() = %q, want %q",
					tt.input, out, tt.expected)
			}
		})
	}
}

// TestTruncateChain tests the chaining capabilities of the Truncate method with other methods
func TestTruncateChain(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     string
		function func(*Conv) *Conv
	}{
		{
			name:  "Uppercase and truncate",
			input: "hello world",
			want:  "HELLO W...",
			function: func(t *Conv) *Conv {
				return t.ToUpper().Truncate(10)
			},
		},
		{
			name:  "Lowercase and truncate",
			input: "HELLO WORLD",
			want:  "hello...",
			function: func(t *Conv) *Conv {
				return t.ToLower().Truncate(8)
			},
		},
		{
			name:  "Truncate and repeat",
			input: "hello",
			want:  "hellohello", // No padding expected before repeat
			function: func(t *Conv) *Conv {
				return t.Truncate(6).Repeat(2)
			},
		},
		{
			name:  "SnakeCase and truncate",
			input: "Hello World Example",
			want:  "hello_world_...",
			function: func(t *Conv) *Conv {
				return t.SnakeLow().Truncate(15)
			},
		},
		{
			name:  "Truncate with custom separator",
			input: "Hello World Example",
			want:  "hello-world-ex...",
			function: func(t *Conv) *Conv {
				return t.SnakeLow("-").Truncate(17)
			},
		},
		{
			name:  "Using explicit reserved chars",
			input: "Hello, World!",
			want:  "Hell...",
			function: func(t *Conv) *Conv {
				return t.Truncate(10, 3)
			},
		},
		{
			name:  "Using different numeric types",
			input: "Testing different types",
			want:  "Testing...",
			function: func(t *Conv) *Conv {
				return t.Truncate(uint8(12), float64(2))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function(Convert(tt.input)).String()
			if got != tt.want {
				t.Fatalf("\n🎯Test: %q\ninput: %q\n   got: %q\n  want: %q", tt.name, tt.input, got, tt.want)
			}
		})
	}
}
