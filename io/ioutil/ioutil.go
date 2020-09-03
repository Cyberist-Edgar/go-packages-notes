package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开一个文件
	file, err := os.Open("data")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败: ", err)
		return
	}

	// 读取file中的全部数据
	data, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		fmt.Println("读取文件失败: ", err1)
	}
	fmt.Println(string(data))

	// ReadFile可以直接指定的文件名中读取数据
	data2, err2 := ioutil.ReadFile("data")
	if err2 != nil {
		fmt.Println("读取文件失败: ", err2)
		return
	}
	fmt.Println(string(data2))

	// WriteFile 可以直接向文件中写入数据
	err3 := ioutil.WriteFile("data2", []byte("Hello this is a test"), 0666)
	if err3 != nil {
		fmt.Println("写入文件失败: ", err3)
		return
	}

	// ReadDir 可以直接获取路径中所有目录和文件的信息
	infoList, err4 := ioutil.ReadDir(".")
	if err4 != nil {
		fmt.Println("读取目录失败: ", err4)
		return
	}

	for _, info := range infoList {
		fmt.Println(info.Name(), info.IsDir())
	}

	// 创建一个临时目录
	name, err5 := ioutil.TempDir(".", "temp")
	if err5 != nil {
		fmt.Println("创建临时目录失败: ", err5)
		return
	}
	fmt.Println("临时目录为: ", name)

	// 创建一个临时文件
	f, err6 := ioutil.TempFile(".", "test")
	if err6 != nil {
		fmt.Println("创建临时文件失败: ", err6)
		return
	}
	defer f.Close()
	f.Write([]byte("Hello this is a temp file"))
	fmt.Println("临时文件名为: ", f.Name())

}
