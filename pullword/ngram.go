package main

import (
	"fmt"
	"strings"
)

func GetNGramFromArray(min, max int, words []string) map[string]int {
	if min <= 0 || max <= 0 || min > max {
		return nil
	}
	dict := make(map[string]int)
	n := len(words)
	for i := 0; i < n; i++ {
		for j := min; j <= max; j++ {
			if i+j > n {
				break
			}
			dict[strings.Join(words[i:i+j], " ")]++
		}
	}
	return dict
}

func main() {
	words := []string{"let", "us"} //, "talk", "about", "what", "will", "do"}
	ret := GetNGramFromArray(1, 2, words)
	fmt.Printf("%+v", ret)
}
