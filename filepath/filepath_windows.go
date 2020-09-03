package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 该代码演示在window条件的输出示例
func main() {
	fmt.Println("path/filepath package的使用")

	var absolute = `D:\Documents\Go\learn\src\packages\path\filepath`
	var relative = `..\..\sync\atomic`

	fmt.Println(filepath.IsAbs(absolute)) // true
	fmt.Println(filepath.IsAbs(relative)) // false

	fmt.Println(filepath.Abs(absolute)) // D:\Documents\Go\learn\src\packages\path\filepath <nil>
	fmt.Println(filepath.Abs(relative)) // D:\Documents\Go\learn\src\packages\sync\atomic <nil>

	splitListExample := `sync\atomic\;path\filepath`
	fmt.Println(filepath.SplitList(splitListExample)) // [sync\atomic\ path\filepath]

	fmt.Println(filepath.Split(absolute)) // D:\Documents\Go\learn\src\packages\path\ filepath

	fmt.Println(filepath.Join(absolute, relative)) // D:\Documents\Go\learn\src\packages\sync\atomic

	fmt.Println(filepath.FromSlash("/a/b/c")) // \a\b\c

	fmt.Println(filepath.ToSlash(`\a\b\c`)) // /a/b/c

	fmt.Println(filepath.VolumeName(`D:\Documents\Go`)) // D:

	fmt.Println(filepath.Dir(`D:\Documents\Go\learn`))

	fmt.Println(filepath.Base(`D:\Documents\Go\learn`))

	fmt.Println(filepath.Ext("data.csv"))

	fmt.Println(filepath.Clean(`path\..\sync\atomic`))

	fmt.Println(filepath.EvalSymlinks(`./atomic_value.go.lnk`))

	fmt.Println(filepath.Match(`D:\*\Go\learn`, `D:\Documents\Go\learn`)) // true <nil>

	fmt.Println(filepath.Glob(`D:\*\Go\learn`)) // [D:\Documents\Go\learn] <nil>

	filepath.Walk(".", walkFunc)
	/*
		README.md README.md false <nil>
		atomic_value.go.lnk atomic_value.go.lnk false <nil>
		filepath.exe filepath.exe false <nil>
		filepath_unix.go filepath_unix.go false <nil>
		filepath_windows.go filepath_windows.go false <nil>
	*/
}
func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path, info.Name(), info.IsDir(), err)
	return err
}
