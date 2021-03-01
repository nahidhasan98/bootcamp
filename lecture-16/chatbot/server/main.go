package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//just a message for ensuring that local server is running
	fmt.Println("Local server is running on port 5000")

	//listening to the port
	listen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//accepting all incoming request & creating session
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	//loop for chatting/request-response
	for {
		//preparing the incoming message
		msg, size := receiveFromClient(conn)

		//printing the incoming msg
		displayMsg(msg, size)

		//preparing the outgoing message
		response := prepareResponse(msg, size)

		//response back to client
		sendToClient(conn, response)
	}
}

func receiveFromClient(conn net.Conn) (string, int) {
	msgIn := make([]byte, 1024)
	size, err := conn.Read(msgIn)
	checkErr(err)
	msg := string(msgIn)
	msg = msg[:size]
	msg = strings.TrimSpace(msg)

	return msg, size
}

func displayMsg(msg string, size int) {
	fmt.Println("Client: " + msg)
}

func prepareResponse(msg string, size int) string {
	var response string

	if strings.ToLower(msg) == "date" {
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		weekDay := time.Now().Weekday()

		response = strconv.Itoa(day) + "-" + month.String()[:3] + "-" + strconv.Itoa(year) + " " + weekDay.String()
	} else if strings.ToLower(msg) == "time" {
		response = time.Now().Format(time.Kitchen)
	} else {
		resp, err := strconv.Atoi(msg)
		checkErr(err)

		if msg != "0" && resp == 0 {
			response = "Sorry, I can't understand what you are saying. :("
		} else {
			resp++
			response = strconv.Itoa(resp)
		}
	}

	//printing the outgoing msg
	fmt.Println("You: " + response)

	return response
}

func sendToClient(conn net.Conn, response string) {
	conn.Write([]byte(response))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
