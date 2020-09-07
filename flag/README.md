# flag 基本使用

## 方法使用

在 flag 里面主要有两种方式定义一个命令行参数:

一种方式是类似`typ(name, value, usage)`，该方式返回一个 `*typ` 类型的数据，该数据保存对应参数的值，另一种方式为`typVar(&typ, name, value, usage)`，该方式会将对应的参数值保存在typ变量里面。在使用命令行参数之前必须要调用`flag.Parse()`方法
```
这里的typ可以为：
Bool, Int, Int64, Uint, Uint64, Float64, String, Duration
```
使用例子：
```go
var name string
// 将name和命令行参数进行绑定
flag.StringVar(&name, "name", "", "your name")

// 当然你也可以使用如下的方式进行绑定
name := flag.String("name", "", "your name")
// 但是你访问绑定的参数值时，应该使用 *
fmt.Println(*name)
```
当然你可以实现Value接口，然后将该参数传入`func Var(value Value, name string, usage string)`中进行绑定
```go
type Value interface {
    String() string
    Set(string) error
}
```

`func Arg(i int) string`

返回解析之后剩余参数中的第 i 个参数，索引从 0 开始，这些参数是没有指定 name 的

`func Args() []string`

Args 返回解析之后剩余的非 flag 参数

`func NArg() int`

NArg 返回 flags 处理之后剩余的参数个数(这些参数不能有指定的参数名，否则会报错)

`func NFlag() int`

NFlag 返回设置的命令行参数的个数(在命令行中提到的)

`func Parse()`

Parse 从 os.Args[1:]中解析命令行参数，必须在所有的参数定义好之后，使用 flag 之前调用

`func Parsed() bool`

Parsed 返回是否命令行参数已经被解析好了

`func PrintDefaults()`

PrintDefaults 会向标准错误(如果没有进行其他配置)中写入所有注册好的 flag 的使用信息，比如说对于 int 类型的 flag x，默认的输入形式如下：

```
-x int
	x的使用介绍 (default 7)
```

`func Set(name, value string) error`

Set 设置已经注册好的 flag 的值

`func UnquoteUsage(flag *Flag) (name string, usage string)`

UnquoteUsage 从一个 flag 的使用方法字符串中提取一个反引号包括的名称，返回该值以及没有包括的使用方法。 比如说输入"a \`name\` to show" 返回 ("name", "a name to show")。如果其中没有反引号，那么会返回有根据猜测的值，如果该 flag 为 bool 类型返回空字符串

`func Visit(fn func(*Flag))`

Visit 以字典顺序访问已经设置好的命令行参数，并且对每个 flag 执行 fn

`func VisitAll(fn func(*Flag))`

VisitAll 以字典顺序访问注册好的命令行参数，并且对每个 flag 执行 fn

## Flag

```go
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```

Flag 代表了一个 flag 的信息

`func Lookup(name string) *Flag`

Lookup 返回命令行参数名为 name 的 Flag 对象，如果不存在，那么返回 nil

## FlagSet

```go
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage func()
	// contains filtered or unexported fields
}
```

FlagSet 代码了一系列被定义好的 flag，FlagSet 的零值没有 name，且包括一个 ContinueOnError 错误处理

Flag 的名称必须和 FlagSet 里面的名称不同，否则会 panic。

FlagSet可以定义命令行，实际上调用flag包定义的参数都是间接通过FlagSet进行设置的，所以你可以类似下面这样的方式来定义一个参数：
```go
set := flag.NewFlagSet("cli", ContinueOnError)
var name string 
set.StringVar(&name, "name", "", "your name")
```


`func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet`

NewFlagSet 使用特定的 name(name 为本程序的命名) 和错误处理属性返回一个新的空 flag 集合。如果 name 不是非空，它会打印在默认用法消息和错误消息

ErrorHandling类型有如下三种：
```go
const (
    ContinueOnError ErrorHandling = iota
    ExitOnError
    PanicOnError
)
```

`func (f *FlagSet) ErrorHandling() ErrorHandling`

ErrorHandling 返回 FlagSet 的错误处理模式

`func (f *FlagSet) Init(name string, errorHandling ErrorHandling)`

Init 方法为 FlagSet 设置参数名以及错误处理。默认条件下，FlagSet 零值使用一个空的 name 和 ContinueOnError 处理策略

`func (f *FlagSet) SetOutput(output io.Writer)`

SetOutPut 设置使用和错误信息的写入的地址，如果output为nil，那么会使用os.Stderr