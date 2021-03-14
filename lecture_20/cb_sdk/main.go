package main

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

var collection *gocb.Collection

func init() {
	// //couchbase connection block
	cluster, err := gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "root",
			Password: "bootcamp",
		})
	checkErr(err)

	// get a bucket reference
	bucket := cluster.Bucket("bootcamp")

	// We wait until the bucket is definitely connected and setup.
	err = bucket.WaitUntilReady(5*time.Second, nil)
	checkErr(err)

	// get a collection reference
	collection = bucket.DefaultCollection()

	// for a named collection and scope
	// scope := bucket.Scope("my-scope")
	// collection := scope.Collection("my-collection")
}

type RequestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func main() {
	// Upsert Document Start //

	//using map data
	// upsertData := map[string]string{
	// 	"name":    "Nahid",
	// 	"company": "Master Academy",
	// 	"email":   "mnh.nahid35@gmail.com",
	// }

	//using struct data
	// var upsertData RequestTable
	// upsertData.Name = "Nahid"
	// upsertData.Company = "Master Academy"
	// upsertData.Email = "mnh.nahid35@gmail.com"
	// upsertData.Type = "request"
	// upsertData.Status = 1

	//using struct literal
	upsertData := RequestTable{
		ID:      "request::3",
		Name:    "Nahid",
		Company: "Master Academy",
		Email:   "mnh.nahid35@gmail.com",
		Type:    "request",
		Status:  1,
	}

	upsertResult, err := collection.Upsert("request::3", upsertData, &gocb.UpsertOptions{})
	checkErr(err)

	fmt.Println(upsertResult.Cas())
	// Upsert Document End //
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
