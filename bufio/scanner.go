package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := `
	Hello World
	This is a test string 
	created by Edgar
	`
	// 创建一个新的Scanner
	scanner := bufio.NewScanner(strings.NewReader(str))

	// 设置分割函数
	scanner.Split(bufio.ScanLines)

	// 读取数据
	for scanner.Scan() {
		test := scanner.Text()
		fmt.Println(test)
	}
}
