package main

import(
	"fmt"
	"bytes"
	"log"
)

var buff bytes.Buffer

func main(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("err: ",err)

			// 输出日志写入的内容
			fmt.Println(buff.String())
		}
	}()

	// 自定义一个日志对象
	// 默认的日志写入到buff中
	myLog := log.New(&buff, "--", log.LstdFlags | log.Lmsgprefix)

	fmt.Println(myLog.Flags())
	fmt.Println(myLog.Prefix())

	// 写入日志
	myLog.Println("line 1")
	myLog.Printf("line %d", 2)

	myLog.Panic("log panic")
	
}