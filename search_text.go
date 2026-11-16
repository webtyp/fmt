package fmt

// ToSearchText turns the text into its search form: lower case, no accents,
// «ñ» as «n»; inside a word the dots, dashes, slashes and «+» are dropped
// ("12.345.678-K" → "12345678k", the check letter is kept); words are
// separated by any other sign; words made only of digits that are separated by
// plain spaces only are joined ("+56 9 1234 5678" → "56912345678"), while a
// sign such as «·» ends the field and keeps what follows apart.
//
// It is the ONE function that writes the searchable column when a record is
// saved and that prepares what the person types, so both always agree:
//
//	Convert("Lucía Pérez · 12.345.678-K · +56 9 1234 5678").ToSearchText().String()
//	// "lucia perez 12345678k 56912345678"
func (t *Conv) ToSearchText() *Conv {
	if t.hasContent(BuffErr) || t.outLen == 0 {
		return t
	}
	src := t.GetString(BuffOut)
	var (
		out        = make([]rune, 0, len(src))
		word       = make([]rune, 0, 16)
		wordDigits = true  // every rune of the word is a digit
		prevDigits = false // the last word written was digits only
		spaceOnly  = true  // only spaces lie between the last word and the next
	)
	flush := func(hard bool) {
		if len(word) > 0 {
			joins := prevDigits && wordDigits && spaceOnly && len(out) > 0
			if len(out) > 0 && !joins {
				out = append(out, ' ')
			}
			out = append(out, word...)
			prevDigits = wordDigits
			word = word[:0]
			wordDigits = true
			spaceOnly = true
		}
		if hard {
			spaceOnly = false
		}
	}
	for _, r := range src {
		c, kind := searchRune(unaccent(r))
		switch kind {
		case runeDropped:
		case runeSpace:
			flush(false)
		case runeSign:
			flush(true)
		default:
			word = append(word, c)
			if !isDigit(c) {
				wordDigits = false
			}
		}
	}
	flush(false)
	t.ResetBuffer(BuffOut)
	t.WrString(BuffOut, string(out))
	return t
}

// What ToSearchText does with one rune.
const (
	runeKept    = iota // part of a word
	runeDropped        // vanishes inside a word
	runeSpace          // separates words, joinable digits stay joinable
	runeSign           // separates words and ends the field
)

// searchRune maps one rune for ToSearchText: the rune to write (lower case,
// ñ as n) and what to do with it.
func searchRune(r rune) (rune, int) {
	switch r {
	case 'ñ', 'Ñ':
		return 'n', runeKept
	case '.', '-', '/', '+':
		return 0, runeDropped
	case ' ', '\t', '\n', '\r':
		return 0, runeSpace
	}
	switch {
	case r >= 'A' && r <= 'Z':
		return r + 32, runeKept
	case r >= 'a' && r <= 'z', isDigit(r):
		return r, runeKept
	case r < 0x80, r <= 0xBF, r >= 0x2000 && r <= 0x206F:
		return 0, runeSign // ASCII, Latin-1 and general punctuation («·», «–», «…»)
	}
	return toLowerRune(r), runeKept
}

func isDigit(r rune) bool { return r >= '0' && r <= '9' }
