package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // create a random

	divideNumber(10, 2)

	randomVal := randomRange(20, 30)
	fmt.Println("random value: ", randomVal)

	var names = []string{"bagus", "merta"}
	printMsg("Hello ", names)

	multiply()

	fmt.Println(calculate(7))

}

func randomRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value // int
}

func printMsg(msg string, arr []string) {
	var name = strings.Join(arr, " ")
	fmt.Println(msg, name) // msg + name
}

func multiply() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return //area, circumference
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divide by %d\n", m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}
