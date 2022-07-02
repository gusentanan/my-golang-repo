package main

import "fmt"

func main() {
	value := 60
	if value == 100 {
		fmt.Println("You Geniuses!")
	} else if value > 80 {
		fmt.Println("You Got A")
	} else if value < 80 && value == 70 {
		fmt.Println("You Got B")
	} else {
		fmt.Println("You DUMB!")
	}

	var point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s Marvelous!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s Good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s Not Bad\n", percent, "%")
	}

	var point2 = 76
	switch {
	case point2 == 100:
		fmt.Println("You're GOAT")
	case (point2 <= 80) && (point2 >= 50):
		fmt.Println("Well at least you try")
		fallthrough // continue
	case point2 < 40:
		fmt.Println("Your better pack your things now!")
	default:
		{
			fmt.Println("Did you take the exam?")
		}
	}

	var point3 = 10
	if point2 > 7 {
		switch point3 {
		case 10:
			fmt.Println("Perfect scores!")
		default:
			fmt.Println("Hmmm")
		}
	} else {
		if point2 == 5 {
			fmt.Println("Not bad")
		} else if point2 == 3 {
			fmt.Println("Neva Give Up")
		} else {
			fmt.Println("Meh")
			if point2 == 0 {
				fmt.Println("you better see me in my office right now kid!")
			}
		}
	}

}
