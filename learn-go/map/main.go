package main

import "fmt"

func main() {
	// Just like HashMap in kotlin
	var extMap map[string]int
	extMap = map[string]int{}
	extMap["Atomic Habits"] = 80
	extMap["Psychology of Money"] = 84

	extMap = map[string]int{
		"Why We Sleep":                   79,
		"The Subtle Art to Not Give a F": 89,
		"Atomic Habits":                  80,
		"Psychology of Money":            84,
	}

	fmt.Println("Atomic Habits Score:", extMap["Atomic Habits"])

	for key, val := range extMap {
		fmt.Println(key, ": ", val)
	}

	// delete
	delete(extMap, "Why We Sleep")
	fmt.Println(len(extMap))
	fmt.Println(extMap)

	// is Exist?
	var value, isExist = extMap["Atomic Habits"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Value not exist!")
	}
}
