package fmt_test

import (
	"testing"

	. "webtyp.com/fmt"
)

func TestRemoveAccentsRemovesSpanishAccents(t *testing.T) {
	cases := map[string]string{
		"Lucía":       "Lucia",
		"Hóla Múndo":  "Hola Mundo",
		"ÁÉÍÓÚ áéíóú": "AEIOU aeiou",
		"Pingüino":    "Pinguino",
		"sin tilde":   "sin tilde",
		"Ñandú":       "Ñandu",
	}
	for in, want := range cases {
		if got := Convert(in).RemoveAccents().String(); got != want {
			t.Errorf("RemoveAccents(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestRemoveAccentsChainsForSearch(t *testing.T) {
	if got := Convert("Hóla Múndo").RemoveAccents().ToLower().String(); got != "hola mundo" {
		t.Errorf("RemoveAccents().ToLower() = %q, want %q", got, "hola mundo")
	}
	if got := Convert("").RemoveAccents().String(); got != "" {
		t.Errorf("empty text = %q", got)
	}
}
