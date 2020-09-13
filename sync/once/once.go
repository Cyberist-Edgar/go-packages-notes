package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var once sync.Once
	for i := 0; i < 100; i++ {
		// once 只会执行一次，尽管多次调用
		once.Do(func() {
			count++
		})
		// once只会指定一次Do中的函数，函数定义不同也不会再次执行
		once.Do(func() {
			fmt.Println("another func but won't excute")
		})
	}
	fmt.Println("count: ", count)
}
