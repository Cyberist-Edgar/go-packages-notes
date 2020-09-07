package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {
	var buff bytes.Buffer
	// 创建一个Writer，最终数据写入到buff中
	writer := multipart.NewWriter(&buff)
	// 创建一个表单数据
	writer.WriteField("key", "value")
	// 创建一个传送文件的表单，第二个参数可以自定义
	w, err := writer.CreateFormFile("img", "demo.png")
	if err != nil {
		fmt.Println("创建文件失败: ", err)
		return
	}
	// 读取文件中的内容
	data, err := ioutil.ReadFile("demo.png")
	if err != nil {
		fmt.Println("读取图片发生错误: ", err)
		return
	}
	// 写入到writer中
	w.Write(data)
	writer.Close()

	// 创建一个post请求
	req, err := http.NewRequest("POST", "http://www.httpbin.org/post", &buff)
	if err != nil {
		fmt.Println("创建请求失败: ", err)
	}

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	// 读取返回的内容
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	fmt.Println(string(d))
}
