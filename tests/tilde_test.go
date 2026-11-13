package fmt_test

import (
	"testing"

	. "webtyp.com/fmt"
)

func TestTildeRemovesSpanishAccents(t *testing.T) {
	cases := map[string]string{
		"Lucía":       "Lucia",
		"Hóla Múndo":  "Hola Mundo",
		"ÁÉÍÓÚ áéíóú": "AEIOU aeiou",
		"Pingüino":    "Pinguino",
		"sin tilde":   "sin tilde",
		"Ñandú":       "Ñandu",
	}
	for in, want := range cases {
		if got := Convert(in).Tilde().String(); got != want {
			t.Errorf("Tilde(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestTildeChainsForSearch(t *testing.T) {
	if got := Convert("Hóla Múndo").Tilde().ToLower().String(); got != "hola mundo" {
		t.Errorf("Tilde().ToLower() = %q, want %q", got, "hola mundo")
	}
	if got := Convert("").Tilde().String(); got != "" {
		t.Errorf("empty text = %q", got)
	}
}
