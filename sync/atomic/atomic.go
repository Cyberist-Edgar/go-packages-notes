package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	value int64 = 0
)

func changeValue(wg *sync.WaitGroup) {
	defer wg.Done()
	// 使用下面的方式对value进行操作不会出现bug
	// 如果只是使用value++，则输出的结果与预期不一致
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&value, 1)
	}
}

func main() {
	fmt.Println(value)
	var wg sync.WaitGroup
	wg.Add(2)
	go changeValue(&wg)
	go changeValue(&wg)
	wg.Wait()
	fmt.Println(value)
}
