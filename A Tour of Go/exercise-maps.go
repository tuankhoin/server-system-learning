package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sub := strings.Fields(s)
	m := make(map[string]int)
	for i := range sub{
		m[sub[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
