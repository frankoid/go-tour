package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count, _ := counts[word]
		counts[word] = count + 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
