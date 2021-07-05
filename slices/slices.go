package main

import (
	"golang.org/x/tour/pic"
)

func Pic(x, y int) [][]uint8 {
	var sl = make([][]uint8, y)

	//alocation
	for i := range sl {
		sl[i] = make([]uint8, x)

	}

	//draw
	for i := range sl {
		for j := range sl[i] {
			sl[i][j] = uint8((i ^ j | i*j)) // play around the numbers a bit
		}
	}
	return sl
}

func main() {
	pic.Show(Pic)
}
