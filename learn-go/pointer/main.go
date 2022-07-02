package main

import (
	"fmt"
)

func main() {

	var number1 int = 4
	var number2 *int = &number1
	fmt.Println("number1 (value)   :", number1)  // 4
	fmt.Println("number1 (address) :", &number1) // 0xc20800a220
	fmt.Println("number2 (value)   :", *number2) // 4
	fmt.Println("number2 (address) :", number2)  // 0xc20800a220
	fmt.Println("")

	number1 = 5
	fmt.Println("number1 (value)   :", number1)  // 5
	fmt.Println("number1 (address) :", &number1) // 0xc20800a220
	fmt.Println("number2 (value)   :", *number2) // 5
	fmt.Println("number2 (address) :", number2)  // 0xc20800a220

	var number3 = 4
	fmt.Println("before", number3) // 4
	changeNumber(&number3, 10)
	fmt.Println("after", number3) // 10
}

func changeNumber(original *int, value int) {
	*original = value
}
