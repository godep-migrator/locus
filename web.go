package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
	"os"
)

func serveHttp() {
	m := martini.Classic()
	m.Use(render.Renderer())

	// UI Calls
	m.Get("/", homeHandler)
	m.NotFound(notFoundHandler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), m); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("home").ParseFiles("views/layout.html", "views/home/index.html"))
	t.ExecuteTemplate(w, "layout", "Home")
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("404").ParseFiles("views/layout.html", "views/404.html"))
	t.ExecuteTemplate(w, "layout", "Whoops!")
}
