package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number", i)

	}

	var i = 0
	for i < 5 {
		fmt.Println("Number", i)
		i++

	}

	i = 0
	for {
		fmt.Println("Number", i)
		i++
		if i == 5 {
			break
		}

	}

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Number", i)

	}

	// segitiga siku terbaliks
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()

	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]\n")
		}

	}
}
