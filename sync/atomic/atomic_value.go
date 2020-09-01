// 这里只是一个Value的简单使用方法
// 你也可以到https://pkg.go.dev/sync/atomic?tab=doc#example-Value-Config查看其他的例子
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	value atomic.Value
)

func changeValue(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	// 使用下面的方式对value进行操作不会出现bug
	// 如果只是使用value++，则输出的结果与预期不一致
	for i := 0; i < 10000; i++ {
		// 使用m来保证不会有其他的读写
		m.Lock()
		v := value.Load().(int)
		value.Store(v + 1)
		m.Unlock()
	}
}

func main() {
	fmt.Println(value)
	value.Store(0)
	var wg sync.WaitGroup

	var m sync.Mutex
	wg.Add(2)
	go changeValue(&wg, &m)
	go changeValue(&wg, &m)
	wg.Wait()
	fmt.Println(value)
}
