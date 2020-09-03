# path/filepath 基本使用

标准库path/filepath中提供了对各个操作系统文件路径的操作函数，一般情况下不要使用path，因为package path 只兼容以`/`为路径分隔符的系统

## 方法使用

`func IsAbs(path string) bool`

该函数以路径为参数，如果该路径为绝对路径，那么返回true，如果不是返回false

`func Abs(path string) (string, error)`
<img src="https://pic2.zhimg.com/80/v2-e2b8b08e8c225d0ec172c35ed9b3cdd1_1440w.png">
分析源码，我们知道如果输入路径时绝对路径，则对路径进行clean(见Clean函数)之后，返回，如果是相对路径，则会加上当前目录使之成为绝对路径，然后返回


`func Rel(basepath, targpath string) (string, error)`

该函数以basepath为基准，返回targpath相对于basepath的相对路径，也就是说如果basepath为`/a`，targpath为`/a/b/c`，那么则会返回`/b/c`，如果两个参数有一个为绝对路径，一个为相对路径，则会返回错误

`func SplitList(path string) []string`

该函数会将path中存在的多个`环境变量`路径进行分割(这些路径由os.PathListSeparator连接)，比如说`/home/linux:/usr/bin`会分割为`[]string{"home/linux", "/usr/bin"}`, unix下默认分隔符为`:`，windows下为`;`

`func Split(path string) (dir, file string)`

Split 函数会以路径中的最后一个os.PathSeparator为节点进行分割，如果没有，dir设置为空字符串，file直接设为path

`func Join(elem ...string) string`

Join 函数将多个路径进行连接，并且进行Clean操作，然后返回


`func FromSlash(path string) string`

FromSlash函数将path中的斜杠（'/'）替换为路径分隔符并返回替换结果，多个斜杠会替换为多个路径分隔符。

`func ToSlash(path string) string`

ToSlash函数将path中的路径分隔符替换为斜杠（'/'）并返回替换结果，多个路径分隔符会替换为多个斜杠。

`func VolumeName(path string) (v string)`

VolumeName函数返回最前面的卷名。如Windows系统里提供参数"C:\foo\bar"会返回"C:"；Unix/linux系统的"\\host\share\foo"会返回"\\host\share"；其他平台会返回""。(文档中如是说，但是其实在源码中只有window下才会返回对应的卷名，其他的均为空字符串)
<img src="https://pic1.zhimg.com/80/v2-68e88f6838e842563e480bcb0751ff15_1440w.png">
<img src="https://pic2.zhimg.com/80/v2-e8da3505fa7327b4f01ae7423e122c39_1440w.png">

`func Dir(path string) string`

Dir函数返回路径所在的目录


`func Base(path string) string`

Base函数会返回路径的最后一个元素，也就是该路径指明的目录名或者文件名

`func Ext(path string) string`

Ext函数返回路径的扩展名，如果没有返回空字符串

`func Clean(path string) string` 

Clean函数会返回与该路径相同的最短路径，在处理中会将比如`a/../b`替换成`b`，会将`.`去掉等等

`func EvalSymlinks(path string) (string, error)`

EvalSymlinks函数会返回path软连接指向的路径

`func Match(pattern, name string) (matched bool, err error)`

Match要求匹配整个name字符串，而不是它的一部分。只有pattern语法错误时，会返回ErrBadPattern。

Windows系统中，不能进行转义：'\\'被视为路径分隔符。
pattern 语法如下：
```
pattern:
	{ term }
term:
	'*'                                  匹配0或多个非路径分隔符的字符
	'?'                                  匹配1个非路径分隔符的字符
	'[' [ '^' ] { character-range } ']'  字符组（必须非空）
	c                                    匹配字符c（c != '*', '?', '\\', '['）
	'\\' c                               匹配字符c
character-range:
	c           匹配字符c（c != '\\', '-', ']'）
	'\\' c      匹配字符c
	lo '-' hi   匹配区间[lo, hi]内的字符
```

`func Glob(pattern string) (matches []string, err error)`

Glob函数返回所有匹配模式匹配字符串pattern的文件或者nil（如果没有匹配的文件）

`func Walk(root string, walkFn WalkFunc) error`

Walk函数会遍历root指定的目录下的文件树，对每一个该文件树中的目录和文件都会调用walkFn，包括root自身。所有访问文件/目录时遇到的错误都会传递给walkFn过滤，Walk函数不会遍历文件树中的符号链接（快捷方式）文件包含的路径

其中 WalkFunc类型为 func(path string, info os.FileInfo, err error) error

示例代码

[filepath_windows.go](filepath_windows.go)
[filepath_unix.go](filepath_unix.go)