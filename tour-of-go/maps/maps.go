package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)
	num := strings.Fields(s)

	for _, word := range num {
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
