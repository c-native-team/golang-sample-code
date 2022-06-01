package main

import (
	"fmt"
	"time"
)

func main() {
	var a string
	for {
		a = input_string("Hello, World")
		fmt.Println(a)
		time.Sleep(5 * time.Second)
	}
}

func input_string(x string) string {
	return x
}
