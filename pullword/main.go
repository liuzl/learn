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

func PointCount(m map[string]*TokenInfo) map[string]*TokenInfo {
	for k, v := range m {
		terms := strings.Fields(k)
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	m := make(map[string]*TokenInfo)
	var total float64
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		ret := goutil.GetNGram(1, 6, line)
		for k, v := range ret {
			if m[k] == nil {
				m[k] = &TokenInfo{}
			}
			m[k].Freq += v
			total += v
		}
	}

}
