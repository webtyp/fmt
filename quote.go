package fmt

// quote wraps the output in double quotes as a Go string literal: the engine of
// Sprintf("%q", …). It escapes " and \, writes \a \b \f \n \r \t \v by name and
// \xHH for every other byte < 0x20 and 0x7f. Bytes >= 0x80 pass through, so
// non-printable non-ASCII runes are NOT escaped, unlike Go's %q: that would need
// Unicode printability tables, too costly in a WASM binary.
func (c *Conv) quote() *Conv {
	if c.hasContent(BuffErr) {
		return c // Error chain interruption
	}
	if c.outLen == 0 {
		c.ResetBuffer(BuffOut)
		c.WrString(BuffOut, quoteStr)
		return c
	}

	// Use work buffer to build quoted string, then swap to output
	c.ResetBuffer(BuffWork)
	c.wrByte(BuffWork, '"')

	// Process buffer directly without string allocation (like capitalizeASCIIOptimized)
	for i := 0; i < c.outLen; i++ {
		char := c.out[i]
		switch char {
		case '"':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, '"')
		case '\\':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, '\\')
		case '\n':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'n')
		case '\r':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'r')
		case '\t':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 't')
		case '\a':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'a')
		case '\b':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'b')
		case '\f':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'f')
		case '\v':
			c.wrByte(BuffWork, '\\')
			c.wrByte(BuffWork, 'v')
		default:
			if char < 0x20 || char == 0x7f {
				c.wrByte(BuffWork, '\\')
				c.wrByte(BuffWork, 'x')
				c.wrByte(BuffWork, hexDigitsLower[char>>4])
				c.wrByte(BuffWork, hexDigitsLower[char&0x0f])
				continue
			}
			c.wrByte(BuffWork, char)
		}
	}

	c.wrByte(BuffWork, '"')
	c.swapBuff(BuffWork, BuffOut)
	return c
}

const hexDigitsLower = "0123456789abcdef"
