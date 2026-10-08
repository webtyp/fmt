package fmt

// Shared constants for maximum code reuse and minimal binary size
const (

	// Common punctuation
	dotStr      = "."
	spaceStr    = " "
	ellipsisStr = "..."
	quoteStr    = "\"\""

	// ASCII case conversion constant
	asciiCaseDiff = 32
	// Buffer capacity constants
	defaultBufCap = 16 // default buffer size
)

// Index-based character mapping for maximum efficiency
var (
	// Accented characters (lowercase) — exported for webtyp/model
	AL = []rune{'á', 'à', 'ã', 'â', 'ä', 'é', 'è', 'ê', 'ë', 'í', 'ì', 'î', 'ï', 'ó', 'ò', 'õ', 'ô', 'ö', 'ú', 'ù', 'û', 'ü', 'ý'}
	// Accented characters (uppercase) — exported for webtyp/model
	AU = []rune{'Á', 'À', 'Ã', 'Â', 'Ä', 'É', 'È', 'Ê', 'Ë', 'Í', 'Ì', 'Î', 'Ï', 'Ó', 'Ò', 'Õ', 'Ô', 'Ö', 'Ú', 'Ù', 'Û', 'Ü', 'Ý'}

	// Private aliases for backward compatibility within fmt
	aL = AL
	aU = AU
)

// toUpperRune converts a single rune to uppercase using optimized lookup
func toUpperRune(r rune) rune {
	// ASCII fast path
	if r >= 'a' && r <= 'z' {
		return r - asciiCaseDiff
	}
	// Accent conversion using index lookup
	for i, char := range aL {
		if r == char {
			return aU[i]
		}
	}
	return r
}

// toLowerRune converts a single rune to lowercase using optimized lookup
func toLowerRune(r rune) rune {
	// ASCII fast path
	if r >= 'A' && r <= 'Z' {
		return r + asciiCaseDiff
	}
	// Accent conversion using index lookup
	for i, char := range aU {
		if r == char {
			return aL[i]
		}
	}
	return r
}

// =============================================================================
// CENTRALIZED WORD SEPARATOR DETECTION - SHARED BY CAPITALIZE AND TRANSLATION
// =============================================================================

// IsWordSeparator checks if a character is a word separator
// UNIFIED FUNCTION: Handles byte, rune, and string inputs in a single function
// OPTIMIZED: Uses IsWordSeparatorChar as single source of truth
func IsWordSeparator(input any) bool {
	switch v := input.(type) {
	case byte:
		return IsWordSeparatorChar(rune(v))
	case rune:
		return IsWordSeparatorChar(v)
	case string:
		// Handle empty strings
		if len(v) == 0 {
			return false
		}
		// Multi-char strings: check if they start with space or newline (translation context)
		if len(v) > 1 && (v[0] == ' ' || v[0] == '\t' || v[0] == '\n') {
			return true
		}
		// Single character strings using the centralized logic
		if len(v) == 1 {
			return IsWordSeparatorChar(rune(v[0]))
		}
		// Check if string ends with newline (separator behavior for translation)
		return v[len(v)-1] == '\n'
	}
	return false
}

// IsWordSeparatorChar is the core separator detection logic
// CENTRALIZED: Single source of truth for what constitutes a word separator
// OPTIMIZED: Handles both ASCII and Unicode characters efficiently
func IsWordSeparatorChar(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' ||
		r == '/' || r == '+' || r == '-' || r == '_' || r == '.' ||
		r == ',' || r == ';' || r == ':' || r == '!' || r == '?' ||
		r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}'
}
