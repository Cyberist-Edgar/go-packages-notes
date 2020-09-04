package main

import (
	"fmt"
	"path"
)

func main() {
	// Base 函数
	fmt.Println(path.Base("/a/b")) // b
	fmt.Println(path.Base("/"))    // /
	fmt.Println(path.Base(""))     // .

	// Clean 函数
	fmt.Println(path.Clean("/a/b/../c"))
	fmt.Println(path.Clean("/a/b/./c"))

	// Dir 函数
	fmt.Println(path.Dir("/a/b/"))     // /a/b
	fmt.Println(path.Dir("/a/b/../c")) // /a
	fmt.Println(path.Dir(""))          // .

	// Ext 函数
	fmt.Println(path.Ext("/a/b"))          // ""
	fmt.Println(path.Ext("/a/b/data.txt")) // .txt

	// IsAbs 函数
	fmt.Println(path.IsAbs("/a/b"))   // true
	fmt.Println(path.IsAbs("../a/b")) // false

	// Join 函数
	fmt.Println(path.Join("a", "b"))     // a/b
	fmt.Println(path.Join("a", "", "b")) // a/b

	// Match 函数
	fmt.Println(path.Match("/a/*/c", "/a/b/c")) // true <nil>
	fmt.Println(path.Match("a*/b", "/a/c/b"))   // false <nil>

	// Split 函数
	fmt.Println(path.Split("/a/b/c"))  // /a/b/ c
	fmt.Println(path.Split("/a/b/c/")) // /a/b/c  ""(空字符串)
}
