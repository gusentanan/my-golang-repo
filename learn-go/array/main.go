package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "I Putu"
	names[1] = "Bagus Merta"
	names[2] = "Sentana"
	fmt.Println(names[0], names[1], names[2])

	var vegetables = [4]string{"carrot", "cabbages", "green leaf", "onion"}
	fmt.Printf("length of the vegetables: %d\n", len(vegetables))
	fmt.Printf("the vegetables: %s\n", vegetables)

	var numbers = [...]int{1, 2, 3, 4, 5, 100}
	fmt.Println("length of the numbers: ", len(numbers))
	fmt.Println("the numbers: ", numbers)

	var nums2 = [2][3]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println("the nums2: ", nums2)

	for i := 0; i < len(vegetables); i++ {
		fmt.Printf("element in vegetables %d: %s\n", i, vegetables[i])
	}

	for i, veget := range vegetables {
		fmt.Printf("The element of %d is %s\n", i, veget)
	}

	for _, veget := range vegetables {
		fmt.Printf("Yoo The vegetables is %s\n", veget)
	}

	var fruits = make([]string, 2)
	fruits[0] = "Durian"
	fruits[1] = "Manggis"
	fmt.Println(fruits)

}
