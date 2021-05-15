package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	return counts
}

func main() {
	for input, expected := range testCases {
		observed := WordCount(input)
		fmt.Printf("%v, expected: %v, observed: %v\n", input, expected, observed)
	}
}

var testCases = map[string]map[string]int{
	"hi there hi":      map[string]int{"hi": 2, "there": 1},
	"what's what what": map[string]int{"what's": 1, "what": 2},
}
