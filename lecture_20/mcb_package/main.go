package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

//mysql table schema/structure
//document 20 MB
type RequstTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func init() {
	db = mcb.Connect("localhost", "root", "bootcamp")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
}

func main() {
	//How to insert into couchbase bucket
	var myData RequstTable

	form := make(url.Values)
	form.Add("bucket", "bootcamp") //bucket Name
	form.Add("aid", "d01")         //document ID
	form.Add("name", "Alex")
	form.Add("company", "Alex LTD")
	form.Add("email", "alex@mail.com")
	form.Add("type", "request") //what type of data or table name in general (SQL)
	form.Add("status", "1")

	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

	// //How to retrieve from couchbase bucket (selected fields only)
	// pres := db.Query("SELECT aid, name, company, email, type, status FROM bootcamp WHERE type = 'request'")
	// rows := pres.GetRows()
	// fmt.Println("Total Rows:", len(rows))
	// fmt.Println(rows)
}
