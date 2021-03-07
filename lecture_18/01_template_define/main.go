package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//A message for diplsying that the local server is running
	fmt.Println("Local server is running on port 8888")

	//serving perspective request
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	//serving file from server to client
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	//localhost running on port 8888
	http.ListenAndServe(":8888", nil)
}

//tpl variable
var tpl *template.Template

//function runs at start
func init() {
	tpl = template.Must(template.ParseGlob("source/*/*"))
}

//function for serving homepage
func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", nil)
}

//function for serving features page
func features(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "features.gohtml", nil)
}

//function for serving docs page
func docs(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "docs.gohtml", nil)
}
