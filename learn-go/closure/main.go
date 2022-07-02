package main

import (
	"fmt"
)

func main() {

	var getMinMax = func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return
	}

	var numbers = []int{2, 3, 4, 5, 6, 7, 8}
	var min, max = getMinMax(numbers)
	fmt.Printf("Max value is %d, Min value is %d\n", max, min)

	var numbers2 = []int{1, 31, 4, 1, 3, 5, 6, 8, 9, 10}
	var newNums = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("Original Nums: ", numbers)
	fmt.Println("Filtered Nums: ", newNums)

	var max2 = 3
	var totalNums, getNums = findMax(numbers2, max2)
	var theNums = getNums()
	fmt.Println("numbers\t:", numbers2)
	fmt.Printf("find \t: %d\n\n", max2)
	fmt.Println("found \t:", totalNums)
	fmt.Println("value \t:", theNums)
}

// closure
func findMax(number []int, max int) (int, func() []int) {
	var res []int
	for _, e := range number {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
