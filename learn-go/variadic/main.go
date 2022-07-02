package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg1 = calculateAgain(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	result := fmt.Sprintf("The result of calculation function is %.2f", avg1)
	fmt.Println(result)

	var numbers = []int{2, 3, 5, 1, 25, 6, 8}
	var avg2 = calculateAgain(numbers...)
	result2 := fmt.Sprintf("The result of calculation function 2 is %.2f", avg2)
	fmt.Println(result2)

	var food = []string{"Nasi Goreng", "Durian Musang Kaing"}
	yourFavoriteFood("Bagus", food...)
}

// variadic function
func yourFavoriteFood(name string, food ...string) {
	var foodies = strings.Join(food, ", ")
	fmt.Printf("Hello my name is %s\n", name)
	fmt.Printf("My favorite food is %s", foodies)
}

func calculateAgain(numbers ...int) float64 {
	var totalValue = 0
	for _, num := range numbers {
		totalValue += num
	}

	average := float64(totalValue) / float64(len(numbers))
	return average
}
