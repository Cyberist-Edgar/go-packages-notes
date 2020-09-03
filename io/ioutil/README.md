# io/ioutil 基本使用

io/ioutil 包中实现了一些 io 工具函数，拿来即用

## 方法使用

`func ReadAll(r io.Reader) ([]byte, error)`

ReadAll 函数从一个 io.Reader 接口参数中一次性读取所有数据，并返回

`func ReadFile(filename string) ([]byte, error)`

ReadFile 函数从 filename 指定的文件中读取数据并返回文件的内容

`func WriteFile(filename string, data []byte, perm os.FileMode) error`

WriteFile 函数向文件 filename 中写入数据，如果文件存在，会清空文件，但不改变权限，如果文件不存在，则以指定的权限创建文件并写入数据

`func ReadDir(dirname string) ([]os.FileInfo, error)`

ReadDir 函数从指定目录中读取所有目录信息

`func TempDir(dir, prefix string) (name string, err error)`

TempDir 函数在 dir 下创建一个新的使用 prefix 为前缀的临时文件夹，如果 dir 为空字符串，那么使用默认临时文件夹

`func TempFile(dir, prefix string) (f *os.File, err error)`

TempDir 函数会在 dir 目录下创建一个新的、使用 prefix 为前缀的临时文件，以读写模式打开该文件并返回 os.File 指针，可以对文件进行写入或者其他的操作

## 示例代码

[ioutil.go](ioutil.go)
