package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 该代码演示在unix条件下的输出示例
func main() {
	fmt.Println("path/filepath package的使用")

	var absolute = `/home/linux/go/src`
	var relative = `../../go`

	fmt.Println(filepath.IsAbs(absolute)) // true
	fmt.Println(filepath.IsAbs(relative)) // false

	fmt.Println(filepath.Abs(absolute)) // /home/linux/go/src <nil>
	fmt.Println(filepath.Abs(relative)) // /home/linux/go <nil>

	splitListExample := `/home/linux:/usr/bin`
	fmt.Println(filepath.SplitList(splitListExample)) // [/home/linux /usr/bin]

	fmt.Println(filepath.Split(absolute)) // /home/linux/go/ src

	fmt.Println(filepath.Join(absolute, relative)) // /home/linux/go

	fmt.Println(filepath.FromSlash(`\a\b\c`)) // \a\b\c

	fmt.Println(filepath.FromSlash(`/a/b/c`)) // /a/b/c

	fmt.Println(filepath.VolumeName(`\host\share\foo`)) //

	fmt.Println(filepath.Dir(`/home/linux/go`)) // /home/linux

	fmt.Println(filepath.Base(`/home/linux/go`)) // go

	fmt.Println(filepath.Ext("data.csv")) // .csv

	fmt.Println(filepath.Clean(`path/../sync/atomic`)) // sync/atomic

	fmt.Println(filepath.EvalSymlinks(`./atomic_value.go.lnk`))

	fmt.Println(filepath.Match(`/home/*/go`, `/home/linux/go`)) // true <nil>

	fmt.Println(filepath.Glob(`/home/*/go`)) // [/home/linux/go] <nil>

	filepath.Walk(".", walkFunc)

}
func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path, info.Name(), info.IsDir(), err)
	return err
}
