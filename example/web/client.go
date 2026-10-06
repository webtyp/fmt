//go:build wasm

// Minimal webtyp/fmt WASM client: it exists to measure the compiled size of fmt,
// so it uses fmt alone (no lang, no other webtyp package).
package main

import (
	"syscall/js"

	. "webtyp.com/fmt"
)

func main() {
	items := []string{"  APPLE  ", "  banana  ", "  piñata  "}

	buf := Convert().Write("<h1>webtyp fmt</h1><ul>")
	for _, item := range items {
		processed := Convert(item).TrimSpace().ToLower().Capitalize().String()
		buf.Write("<li>").Write(processed).Write("</li>")
	}
	buf.Write("</ul>")

	doc := js.Global().Get("document")
	div := doc.Call("createElement", "div")
	div.Set("innerHTML", buf.String())
	doc.Get("body").Call("appendChild", div)

	js.Global().Get("console").Call("log", Sprintf("hello webtyp: %d %.2f %t", 123, 45.67, true))

	select {}
}
