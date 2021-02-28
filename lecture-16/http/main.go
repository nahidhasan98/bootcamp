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
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		//preparing the request body
		reqBody := make([]byte, 1024)

		n, err := conn.Read(reqBody)
		if err != nil {
			fmt.Println(err.Error())
		}

		//printing the request details
		fmt.Println(n)
		fmt.Println(string(reqBody))

		//preparing data for response
		respBody := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>First Server</title>
		</head>
		<body>
			<h1>My Second Server</h1>
			<p>Hello there,<br>
			Welcome you all.</p>
		</body>
		</html>`

		respHeader := "HTTP/1.1 200 OK\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n"
		respHeader = fmt.Sprintf(respHeader, len(respBody))

		respFull := respHeader + respBody

		//printing response details
		fmt.Println(respFull)

		//sending response back
		conn.Write([]byte(respFull))
	}
}
