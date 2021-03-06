package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Local server is running on port 8888")

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	resp := `<h1 style="color:#0099ff;"><center>Hompage</center></h1><br>
	<h1><center>This is hompage</center></h1>`

	fmt.Fprintf(w, resp)
}
func about(w http.ResponseWriter, r *http.Request) {
	resp := `<h1 style="color:#009933;"><center>About</center></h1><br>
	<h1><center>This is about page</center></h1>`

	fmt.Fprintf(w, resp)
}
func contact(w http.ResponseWriter, r *http.Request) {
	resp := `<h1 style="color:#e60000;"><center>Contact</center></h1><br>
	<h1><center>This is contact page</center></h1>`

	fmt.Fprintf(w, resp)
}
