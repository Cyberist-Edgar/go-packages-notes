package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var buff bytes.Buffer
	// 构造一个Writer
	w := bufio.NewWriter(&buff)
	// 一次写入一个字符
	w.Write([]byte{'h'})
	// 一次写入一个字符
	w.WriteByte('e')
	// 写入字符串
	w.WriteString("llo")

	// 将数据写入到底层的io.Writer
	w.Flush()

	fmt.Println(buff.String())
}
