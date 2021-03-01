package main

import (
	"fmt"
	"net"
)

func portscanner(ip string, port string) {
	fmt.Println("Enter target ip address:")
	fmt.Scanln(&ip)
	fmt.Println("Enter target port number:")
	fmt.Scanln(&port)
	_, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Printf("port %s is closed", port)
	} else {
		fmt.Printf("port %s is open", port)
	}

	//conn.Write([]byte("Hello"))
}
func main() {
	ip := "10.10.10.10"
	port := "80"
	portscanner(ip, port)
}
