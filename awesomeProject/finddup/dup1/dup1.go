package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //此处使用map来存储元素，key是string类型，value是int类型
	//make是内置函数，用来新建map
	input := bufio.NewScanner(os.Stdin)
	//bufio包，使用它可以简便和高效地处理输入和输出，其中最有用的一个特性是称为扫描器的类型
	//它可以读取输入，以行或者单词为单位断开，这是处理以行为单位的输入内容的最简单的方式
	for input.Scan() {
		//每次dup从输入中读取一行内容，这一行就作为map中的key，对应的值+1
		//下述counts[input.Text()]++相当于以下两个语句：
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
		if len(counts) > 3 {
			break
		}
	}
	//注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		//此处n表示输入的字符之类，line表示该字符在map中的个数
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
