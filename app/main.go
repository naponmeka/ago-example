package main

import (
	"syscall/js"

	"github.com/naponmeka/ago"
)

func main() {
	println("WASM Go Initialized.")
	doc := js.Global().Get("document")
	doc.Set("title", "Ago example")
	ul := ago.CreateElement("ul", nil, []ago.Element{
		ago.CreateElement("li", nil, ago.CreateElementContent("Gopher")),
		ago.CreateElement("li", nil, ago.CreateElementContent("Dog")),
		ago.CreateElement("li", nil, ago.CreateElementContent("Cat")),
	})
	contentElem := ago.CreateElement("div", nil, []ago.Element{
		ago.CreateElement(
			"p",
			map[string]interface{}{
				"style": map[string]string{
					"color":      "red",
					"font-style": "italic",
				},
				"class": "my-class",
			},
			ago.CreateElementContent("Starting world"),
		),
		ul,
		ago.CreateElement("p", nil, ago.CreateElementContent("end")),
	})
	ago.Render(
		contentElem,
		doc.Call("getElementById", "root"),
	)

}
