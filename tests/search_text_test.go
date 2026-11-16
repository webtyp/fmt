package fmt_test

import (
	"testing"

	. "webtyp.com/fmt"
)

func TestToSearchText(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"name and RUT", "Lucía Pérez · 12.345.678-5", "lucia perez 123456785"},
		{"check letter kept", "12.345.678-K", "12345678k"},
		{"phone words joined", "+56 9 1234 5678", "56912345678"},
		{"full record", "Lucía Pérez · 12.345.678-K · +56 9 1234 5678", "lucia perez 12345678k 56912345678"},
		{"D5 example", "Lucía Pérez · 13047980-0 · +56912345678", "lucia perez 130479800 56912345678"},
		{"enye", "Muñoz Peña", "munoz pena"},
		{"upper enye", "NUÑEZ", "nunez"},
		{"digits then word stay apart", "lucia 1234", "lucia 1234"},
		{"word then digits stay apart", "1234 lucia", "1234 lucia"},
		{"extra spaces", "  Ana   María  ", "ana maria"},
		{"slash in word", "AB/12", "ab12"},
		{"email splits on at", "juan@x.cl", "juan xcl"},
		{"only signs", " · - . ", ""},
		{"empty", "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Convert(c.in).ToSearchText().String(); got != c.want {
				t.Errorf("ToSearchText(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

// What is typed and what is stored go through the same function, so applying
// it twice changes nothing.
func TestToSearchTextIsIdempotent(t *testing.T) {
	in := "Lucía Pérez · 12.345.678-K · +56 9 1234 5678"
	once := Convert(in).ToSearchText().String()
	if twice := Convert(once).ToSearchText().String(); twice != once {
		t.Errorf("not idempotent: %q then %q", once, twice)
	}
}
