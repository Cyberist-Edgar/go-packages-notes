package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := bufio.NewReader(strings.NewReader("Hello world"))
	// 声明一个变量接收
	var p = make([]byte, 1)
	for {
		// 读取数据
		_, err := r.Read(p)
		if err != nil {
			// 文件读取完成
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			fmt.Println("读取出现错误: ", err)
		}
		// 输出读取到的数据
		fmt.Println(string(p))
	}

	b := bufio.NewReader(strings.NewReader("Hello world"))
	// 使用ReadString读取数据
	fmt.Println(r.ReadString('\n'))
}
