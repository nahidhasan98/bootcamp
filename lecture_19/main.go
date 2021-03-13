package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:bootcamp@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()
	fmt.Println("db connection successful")
}
func main() {
	//A message for diplsying that the local server is running
	fmt.Println("Local server is running on port 8888")

	//serving perspective request
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/request", request)

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

func request(w http.ResponseWriter, r *http.Request) {
	//method-1
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	// fmt.Println(name, company, email)
	// fmt.Fprintf(w, `received %s %s %s`, name, company, email) //response

	//method-2
	// r.ParseForm()
	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }

	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
	sql := fmt.Sprintf(qs, name, company, email)
	//fmt.Println(sql)

	insert, err := db.Query(sql)
	checkErr(err)
	defer insert.Close()

	fmt.Fprintf(w, `OK`)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
