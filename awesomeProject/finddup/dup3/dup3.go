package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
ReadFile函数返回一个可以转化成字符串的字节slice,这样它可以被strings.Split分割。
*/

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
