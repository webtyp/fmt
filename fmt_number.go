package fmt

// =============================================================================
// FORMAT NUMBER OPERATIONS - Number formatting with separators and display
// =============================================================================

// isNumericString checks if a string represents a valid number
// Universal helper method - follows buffer API architecture
func (c *Conv) isNumericString(str string) bool {
	if len(str) == 0 {
		return false
	}

	i := 0
	// Handle sign
	if str[0] == '-' || str[0] == '+' {
		i = 1
		if i >= len(str) {
			return false // Just a sign is not a number
		}
	}

	hasDigit := false
	hasDecimal := false

	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			hasDigit = true
		} else if str[i] == '.' && !hasDecimal {
			hasDecimal = true
		} else {
			return false // Invalid character
		}
	}

	return hasDigit // Must have at least one digit
}
