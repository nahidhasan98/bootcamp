package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//just a message for ensuring that local server is running
	fmt.Println("Client app running...")

	//dialling to the port
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	//printing an initial message for client
	fmt.Println("Start chatting:")

	//loop fot chatting
	for {
		//taking user's message as input
		msg := inputUserMsg()

		//sending to the server
		sendToServer(conn, msg)

		//preparing incoming message
		response, size := receiveFromServer(conn)

		//printing message got from server
		displayMsg(response, size)
	}
}

func inputUserMsg() string {
	fmt.Print("You: ")
	var msg string
	fmt.Scan(&msg)

	return msg
}

func sendToServer(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
}

func receiveFromServer(conn net.Conn) (string, int) {
	msgIn := make([]byte, 1024)
	size, err := conn.Read(msgIn)
	checkErr(err)
	response := string(msgIn)

	return response, size
}

func displayMsg(response string, size int) {
	fmt.Println("Bot: " + response[:size])
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
