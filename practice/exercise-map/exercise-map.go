package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount -
func WordCount(s string) map[string]int {
	var resultMap = make(map[string]int)
	charArr := strings.Fields(s)
	for _, v := range charArr {
		if count, ok := resultMap[v]; ok {
			resultMap[v] = count + 1
		} else {
			resultMap[v] = 1
		}
	}
	return resultMap
}

func main() {
	wc.Test(WordCount)
}
