package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/中/?page=1&name=golang#x/y%2Fz")
	if err != nil {
		fmt.Println("解析错误: ", err)
		return
	}
	// 因为u实现了String方法，所以打印出来的结果就是
	// u.String方法返回的结果
	fmt.Println(u) // http://example.com/%E4%B8%AD/?page=1&name=golang#x/y%2Fz
	// 输出锚点的编码形式
	fmt.Println(u.EscapedFragment()) // x/y%2Fz

	// 输出路径的编码形式
	fmt.Println(u.EscapedPath()) // /%E4%B8%AD/
	// 输出路径的非编码形式
	fmt.Println(u.Path) // /中/

	// 输出路径的主机名
	fmt.Println(u.Hostname()) // example.com

	// 判断路径是否为绝对路径
	fmt.Println(u.IsAbs()) // true

	// 类似于String方法，但是输出的结果类型为[]byte
	fmt.Println(u.MarshalBinary())

	// 输出路径主机名的端口号
	fmt.Println(u.Port()) // ""

	// 输出查询数据
	fmt.Println(u.Query()) // map[name:[golang] page:[1]]

	// 返回请求路径
	fmt.Println(u.RequestURI()) // /%E4%B8%AD/?page=1&name=golang

	// 输出路径相对于u的绝对路径
	fmt.Println(u.ResolveReference(&url.URL{Path: "../test"})) // http://example.com/test

	// 会改变原来的URL
	// 该方法类似Parse函数，但是接收参数类型不一致
	fmt.Println(u.UnmarshalBinary([]byte("https://example.org/foo")))
	fmt.Println(u)
}
