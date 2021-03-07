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

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/base.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/base.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/features.gohtml")
	checkErr(err)

	tmpl.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/base.gohtml")
	checkErr(err)

	tmpl, err = tmpl.ParseFiles("wpage/docs.gohtml")
	checkErr(err)

	// tmpl, err = tmpl.ParseFiles("wpage/test.gohtml")
	// checkErr(err)

	tmpl.Execute(w, nil)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
