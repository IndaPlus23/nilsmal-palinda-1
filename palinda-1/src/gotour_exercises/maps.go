package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	return_map := make(map[string]int)

	for index := range words {
		return_map[words[index]] += 1
	}

	return return_map
}

func main_() {
	wc.Test(WordCount)
}
