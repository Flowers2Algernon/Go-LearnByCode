package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
函数os.Open返回两个值，第一个是打开的文件(*os.File),该文件随后被Scanner读取
第二个返回值是一个内置的error类型的值，如果err等于特殊的内置nil值，表示文件成功打开。
文件在被读到结尾的时候，close函数关闭文件，然后释放相应的资源--比如内存
另一方面，如果err不是nil,则说明出错了，这是error的值描述错误原因
简单的错误处理是使用Fprintf和%v在标准错误流上输出一条消息，%v可以使用默认格式显示任意类型的值
此处countLines的调用出现在其声明之前，函数和其他包级别的实体可以以任意次序声明
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if len(counts) > 3 {
			break
		}
	}
}
