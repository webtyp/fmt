package fmt

// RemoveAccents removes the accents of the Spanish vowels (á é í ó ú ü, both
// cases) and returns the text without them: "Lucía" → "Lucia". It is the
// accent-insensitive form used to compare names, not a transliteration: letters
// outside the table (ñ, ç) are kept. Chainable like the other Conv methods:
//
//	Convert("Hóla Múndo").RemoveAccents().ToLower().String() // "hola mundo"
func (t *Conv) RemoveAccents() *Conv {
	if t.hasContent(BuffErr) || t.outLen == 0 {
		return t
	}
	src := t.GetString(BuffOut)
	out := make([]rune, 0, len(src))
	for _, r := range src {
		out = append(out, unaccent(r))
	}
	t.ResetBuffer(BuffOut)
	t.WrString(BuffOut, string(out))
	return t
}

// bL and bU are the base letters of aL and aU, index for index.
var (
	bL = []rune{'a', 'a', 'a', 'a', 'a', 'e', 'e', 'e', 'e', 'i', 'i', 'i', 'i', 'o', 'o', 'o', 'o', 'o', 'u', 'u', 'u', 'u', 'y'}
	bU = []rune{'A', 'A', 'A', 'A', 'A', 'E', 'E', 'E', 'E', 'I', 'I', 'I', 'I', 'O', 'O', 'O', 'O', 'O', 'U', 'U', 'U', 'U', 'Y'}
)

// unaccent maps one rune to its unaccented base; any other rune is returned as is.
func unaccent(r rune) rune {
	for i, a := range aL {
		if r == a {
			return bL[i]
		}
	}
	for i, a := range aU {
		if r == a {
			return bU[i]
		}
	}
	return r
}
