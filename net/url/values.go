package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 解析参数
	values, err := url.ParseQuery("x=1&y=2&y=2;z")
	if err != nil {
		fmt.Println("解析查询参数失败: ", err)
		return
	}
	// 获取x对应的值
	fmt.Println(values.Get("x"))
	// 仅会获取一位
	fmt.Println(values.Get("y"))
	// 会获取所有的值
	fmt.Println(values["y"])

	// 添加一个键值对
	values.Add("key", "value")
	// 将数据进行编码
	fmt.Println(values.Encode())

	// 删除键值对
	values.Del("z")
	values.Del("key")
	fmt.Println(values.Encode())

	values.Set("x", "2")
	fmt.Println(values.Encode())

}
