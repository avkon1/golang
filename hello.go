package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")
}

func Hello(name string) string {
	return "Hello " + name
}
