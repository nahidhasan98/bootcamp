package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//net for network
	//start to listening the request
	nl, err := net.Listen("tcp", ":8888") //ip:port 127.0.0.1/localhost
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1 for exit with unexpected error. 0 for normal exit.
	}

	//Accept al incoming request
	conn, err := nl.Accept() //Layer-5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	//printing the address who sent request
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr) //[::1]:52451

	//8 bit = 1 byte
	//1024 byte = 1 kilo byte
	//map, slice
	bs := make([]byte, 1024) //capacity 1500 bytes. Taking 1024, rest of them are some header fields data
	conn.Read(bs)            //reading the request

	//90 10 20 40 50 70 90
	// H e l l 0
	fmt.Println(string(bs)) //human readable formated message

	//byte
	conn.Write([]byte("welcome to our coding boot camp")) //sending back the respons to the client

	//closing the connection
	conn.Close() //red button press
	nl.Close()   //red button press

}
