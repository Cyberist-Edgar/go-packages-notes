package main 

import (
	"fmt"
	"container/list"
)

func main(){
	// 创建一个链表
	l := list.New()
	// 插入数据，放在链表最后位置
	l.PushBack(1)
	fmt.Println("After push 1, the length of l: ", l.Len())

	// 插入数据2，放在链表最后的位置
	e2 := l.PushBack(2)
	// 插入数据3，放在链表的第一个位置
	e3 := l.PushFront(3)
	traverse(l)

	// 将4插入到e2的前面
	l.InsertBefore(4, e2)
	// 将5插入到e3的后面
	l.InsertAfter(5, e3)
	traverse(l)

	// 移除掉e3
	l.Remove(e3)
	traverse(l)
	
}

func traverse(l *list.List){
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}