package main

import (
	"syscall/js"
	"time"

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
		PageTitle: "My TODO list1",
		Todos: []Todo{
			{Title: "Task 0", Done: false},
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
		},
	}
	agoComponent := ago.CreateComponent(gox, state)
	agoComponent.Render("root")
	state2 := TodoPageData{
		PageTitle: "My TODO list2",
		Todos: []Todo{
			{Title: "Task 0", Done: false},
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
		},
	}
	time.Sleep(2 * time.Second)
	agoComponent.ChangeState(state2)
	state3 := TodoPageData{
		PageTitle: "My TODO list3",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
			{Title: "Task 3", Done: false},
			{Title: "Task 4", Done: false},
			{Title: "Task 5", Done: false},
			{Title: "Task 6", Done: false},
			{Title: "Task 7", Done: false},
		},
	}
	time.Sleep(5 * time.Second)
	agoComponent.ChangeState(state3)
	state4 := TodoPageData{
		PageTitle: "My TODO list4",
		Todos: []Todo{
			{Title: "Task x", Done: false},
			{Title: "Task y", Done: false},
			{Title: "Task z", Done: false},
		},
	}
	time.Sleep(5 * time.Second)
	agoComponent.ChangeState(state4)
	state5 := TodoPageData{
		PageTitle: "My TODO list5",
		Todos: []Todo{
			{Title: "Task x", Done: false},
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: false},
			{Title: "Task 3", Done: false},
		},
	}
	time.Sleep(5 * time.Second)
	agoComponent.ChangeState(state5)

}
