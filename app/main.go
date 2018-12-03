package main

import (
	"syscall/js"

	"github.com/naponmeka/ago"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	println("WASM Go Initialized.")
	doc := js.Global().Get("document")
	doc.Set("title", "Ago example")
	gox := `
	<div>
		<h1>{{.PageTitle}}</h1>
		<ul>
			{{range .Todos}}
				{{if .Done}}
					<li class="done">{{.Title}}</li>
				{{else}}
					<li>{{.Title}}</li>
				{{end}}
			{{end}}
		</ul>
	</div>
	`

	state := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 0", Done: false},
			{Title: "Task 1", Done: true},
			{Title: "Task 2", Done: true},
		},
	}
	agoComponent := ago.CreateComponent(gox, state)
	state2 := TodoPageData{
		PageTitle: "My TODO list2",
		Todos: []Todo{
			{Title: "Task 0", Done: false},
			{Title: "Task 1", Done: true},
			{Title: "Task 2", Done: true},
		},
	}
	agoComponent.ChangeState(state2)

	ago.Render(
		agoComponent.VDom,
		doc.Call("getElementById", "root"),
	)

	// element := ago.Transform(gox, state)

}
