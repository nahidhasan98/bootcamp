package main

import "fmt"

func main() {
	// iteration - for
	var i int

	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 5; j < 11; j++ {
		fmt.Println(j)
	}

	// iteration - while
	// go doesn't provide while syntax. we can use for
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
}
