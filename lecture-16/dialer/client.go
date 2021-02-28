package main

import (
	"fmt"
	"net"
)

func main() {

	//conn, err := net.Dial("tcp", "localhost:8888")
	conn, err := net.Dial("tcp", "91.205.173.170")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	conn.Write([]byte("Hello from Nahid Hasan."))

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bs[:n]))
}
