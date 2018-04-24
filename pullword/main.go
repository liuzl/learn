package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/goutil"
	"io"
	"os"
	"strings"
)

var (
	input = flag.String("i", "input.txt", "input file")
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

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	m := make(map[string]int)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		ret := goutil.GetNGram(1, 6, line)
		for k, v := range ret {
			m[k] += v
		}
	}
	for k, v := range m {
		fmt.Printf("[%s] [%s]: %d\n", k, Reverse(k), v)
	}
}
