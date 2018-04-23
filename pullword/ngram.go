package main

import (
	"fmt"
	"strings"
)

func Reverse(input string) string {
	s := strings.Fields(input)
	return strings.Join(reverse(s), " ")
}

func reverse(ss []string) []string {
	l := len(ss)
	if l <= 1 {
		return ss
	}
	for i := 0; i < l/2; i++ {
		ss[i], ss[l-1-i] = ss[l-1-i], ss[i]
	}
	return ss
}

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
	for k, v := range ret {
		fmt.Printf("%s,%s: %d\n", k, Reverse(k), v)
	}
}
