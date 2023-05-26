package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	fmt.Println("Running at the 5000 port")
	log.Fatal(http.ListenAndServe(":5000", nil))
}