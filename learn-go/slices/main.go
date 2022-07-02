package main

import "fmt"

func main() {
	var buah = []string{"durian", "apel", "melon", "semangka"}
	fmt.Println(buah[0])

	var newBuah = buah[0:2] // first index until 2 index element
	fmt.Println(newBuah)
	fmt.Println(buah[:])  // all of them
	fmt.Println(buah[2:]) // start after index 2  until last element
	fmt.Println(buah[:2]) // first 2 element

	fmt.Println(len(buah))    // 4
	fmt.Println(cap(buah))    // 4
	fmt.Println(len(newBuah)) // 2
	fmt.Println(cap(newBuah)) // 4

	// append
	var anotherNewBuah = append(buah, "Stlobeli")
	fmt.Println(anotherNewBuah)

	// copy
	copied := make([]string, 3)
	actual := []string{"babi", "anjing", "ayam", "Dog"}
	n := copy(copied, actual)
	fmt.Println(copied) // babim anjing, ayam
	fmt.Println(actual) // babi, anjing, ayam, Dog
	fmt.Println(n)      // 3

}
