# log 基本使用

## Logger

```go
type Logger struct {
	// contains filtered or unexported fields
}
```

Logger 表示一个活动状态的记录日志的对象，它会生成一行行的输出写入 io.Writer 接口中。每一条日志操作都会单独调用一次 Writer 的 Write 方法。 Logger 可以被多线程安全的使用

`func New(out io.Writer, prefix string, flag int) *Logger`

New 方法定义一个新的 Logger。out 参数设置日志信息写入的地址，prefix 参数默认会写入到每一行的行首，如果 flag 中定义了 Lmsgprefix ，那么它会写入到日志头后面，日志正文前面，flag 参数定义了日志的属性。

`func (l *Logger) Fatal(v ...interface{})`

Fatal 等价于调用 l.Print()，然后调用 os.Exit(1)退出程序

`func (l *Logger) Fatalf(format string, v ...interface{})`

Fatalf 等价于调用 l.Printf()，然后调用 os.Exit(1)

`func (l *Logger) Fatalln(v ...interface{})`

Fatalln 等价于调用 l.Println()，然后调用 os.Exit(1)

`func (l *Logger) Flags() int`

该方法返回 l 的输出标志 flag, flag 为下面常量中的：

```go
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

`func (l *Logger) Output(calldepth int, s string) error`

Output 将输出写入一次日志事件。s 包含的文本会输出到 logger 中定义好的 prefix 后面。如果 s 最后一个字符不是换行符，那么 s 最后会添加一个换行符。 calldepth 用来恢复 PC，出于一般性考虑而提供，目前预定义的都为 2

`func (l *Logger) Panic(v ...interface{})`

Panic 相当于调用 l.Print()后调用 panic()

`func (l *Logger) Panicf(format string, v ...interface{})`

Panicf 相当于调用 l.Printf()之后调用 panic()

`func (l *Logger) Prefix() string`

Prefix 返回 logger 定义的 prefix

`func (l *Logger) Print(v ...interface{})`

Print 函数调用 l.Output 打印输出到日志中，类似 fmt.Print

`func (l *Logger) Printf(format string, v ...interface{})`

Printf 调用 l.Output 打印输出到日志中，类似 fmt.Printf

`func (l *Logger) Println(v ...interface{})`

Println 调用 l.Output 打印输出到日志中，类似 fmt.Println

`func (l *Logger) SetFlags(flag int)`

SetFlags 设置日志的输出标志

`func (l *Logger) SetOutput(w io.Writer)`

SetOutput 设置了日志输出的地址

`func (l *Logger) SetPrefix(prefix string)`

SetPrefix 设置了日志的前缀

`func (l *Logger) Writer() io.Writer`

Writer 返回了 logger 的输出地址，为 io.Writer 接口类型

在 package log 中预定义了一个 Logger，如下：

```go
var std = New(os.Stderr, "", LstdFlags)
```

默认情况下直接调用log中的方法，我们输出的日志全部写到 os.Stderr 中，并且 flag 定义为 LstdFlags.我们如果不需要自定义日志对象，直接就可以调用方法写入日志即可，比如:

```go
log.Println("Hello World")
```

## 代码示例
[log](log.go)