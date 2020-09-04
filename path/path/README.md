# path 基本使用

package path 实现了以`/`为分隔符路径的相关操作

这个包应该只使用在以`/`为分割符的路径，比如说URLs中的路径，该包不适用于Windows路径，操作处理系统路径使用`path/filepath`包

## 方法使用

`func Base(path string) string`

Base 函数返回path的最后一个函数，在提取最后一个元素之前，末尾的`/`将会被删除， 如果path是空的，那么Base函数返回`.`,如果path只由 / 组成，那么Base返回"/"

`func Clean(path string) string`

Clean 函数会返回等同于path的最短路径，最常见的处理就是`..`的替换和`.`的省略

`func Dir(path string) string`

Dir返回除了最后一个元素以外的所有的元素，一般来说就是path的目录，该函数会对处理之后的路径进行Clean，并移除末尾的`/`，如果path为空字符，则返回`.`

`func Ext(path string) string`

Ext返回路径中的扩展名

`func IsAbs(path string) bool`

IsAbs 返回一个路径是否是绝对路径

`func Join(elem ...string) string`

Join 函数拼接多个路径元素为一个路径，并且处理的结果会进行Clean

`func Match(pattern, name string) (matched bool, err error)`

Match 判断路径是否匹配某个模式，模式语法如下：
```
pattern:
	{ term }
term:
	'*'         matches any sequence of non-/ characters
	'?'         matches any single non-/ character
	'[' [ '^' ] { character-range } ']'
	            character class (must be non-empty)
	c           matches character c (c != '*', '?', '\\', '[')
	'\\' c      matches character c

character-range:
	c           matches character c (c != '\\', '-', ']')
	'\\' c      matches character c
	lo '-' hi   matches character c for lo <= c <= hi
```
如果patter错误，会返回ErrBadPattern

`func Split(path string) (dir, file string)`

Split 函数将一个路径拆分成目录和文件名

## 代码示例
[path.go](path.go)