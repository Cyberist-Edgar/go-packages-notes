package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data2.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"id", "name", "age"})
	writer.WriteAll([][]string{{"1", "Tom", "18"}, {"2", "Jack", "19"}, {"3", "Letty", "20"}})

	if err := writer.Error(); err != nil {
		fmt.Println("执行中发生错误: ", err)
	}
	// 可以调用Flush函数立刻将缓冲区中的数据写入文件中
	writer.Flush()
}
