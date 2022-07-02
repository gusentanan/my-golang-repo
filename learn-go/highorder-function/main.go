package main

import (
	"fmt"
	"strings"
)

type FilterCallback func(string) bool

func main() {
	var data = []string{"Superman", "Batman", "Ironman", "Naruto", "Baby"}
	fmt.Println("Real data: \t", data)

	var dataContainsO = filter1(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	fmt.Println("data that contains o is ", dataContainsO)

	var dataContainslength = filter2(data, func(each string) bool {
		return len(each) == 4
	})
	fmt.Println("data with length of string == 4 is ", dataContainslength)

}

func filter1(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

func filter2(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
