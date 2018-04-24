package main

import (
	"fmt"
	"github.com/liuzl/ling"
	"strings"
)

var nlp = ling.MustNLP(ling.Norm)

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

func GetNGramFromStr(min, max int, input string) map[string]int {
	d := ling.NewDocument(input)
	err := nlp.Annotate(d)
	if err != nil {
		return nil
	}
	return GetNGramFromArray(min, max, d.XRealTokens(ling.Norm))
}

func main() {
	strs := []string{
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，92km到北京`,
		`This is Zhanliang's book. the boys' books.`,
	}
	for _, str := range strs {
		ret := GetNGramFromStr(1, 2, str)
		for k, v := range ret {
			fmt.Printf("[%s] [%s]: %d\n", k, Reverse(k), v)
		}
	}
}
