package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//just a message for ensuring that local server is running
	fmt.Println("Local Server is running on port 8888")

	//listening to the port
	listen, err := net.Listen("tcp", "localhost:8888")
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

	for {
		//preparing the request body
		reqBody := make([]byte, 1024)

		n, err := conn.Read(reqBody)
		if err != nil {
			fmt.Println(err.Error())
		}

		//printing the request details
		fmt.Println(n)
		fmt.Println(string(reqBody)[:n])
	}
}
