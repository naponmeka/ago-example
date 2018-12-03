package main

import (
	"bytes"
	"fmt"
	"html/template"
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
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)
	htmlStr := buf.String()
	fmt.Println(htmlStr)
}
