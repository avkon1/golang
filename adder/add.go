package task2

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Golnag")
}

func Add(number1, number2 int) int {
	return number1 + number2
}

func Repeat(in string, times int) string {

	var ret string

	for i := 1; i < times; i++ {
		ret += ret + in
	}
	return ret
}
