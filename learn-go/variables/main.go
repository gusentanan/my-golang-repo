package main

import "fmt"

func main() {
	var firsName1 string = "Bagus"
	lastName1 := "Merta"
	fmt.Printf("Hello %s %s\n", firsName1, lastName1)

	var firstName2 = "Sentana"
	fmt.Println(firstName2)

	var firstName3, lastName3 = "I Putu Bagus", "Merta Sentana"
	fmt.Println("Hello", firstName3, lastName3+"!")

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

	expc := true
	fmt.Printf("is it true? %t\n", expc)

	msg := `Hi, My name is "I Putu Bagus Merta Sentana"
	Im final year student at Universitas Udayana.`
	fmt.Println(msg)

	var _ = "nothing"
	const nums = 30 //cannot diganti
	fmt.Println(nums)

}
