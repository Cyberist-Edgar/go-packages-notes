# bufio 基本使用

package bufio 实现了缓冲级的 I/O。它包装了 io.Reader 和 io.Writer 结构体对象，创建了 Reader 和 Writer 对象分别实现了这两个接口，但是 Reader 和 Writer 还提供了缓存和其他文本 I/O 操作的方法

## 方法使用

`func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanBytes 是 Scanner 的一个分割函数，每一次返回一个字节作为标记

`func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanLines 是 Scanner 的一个分割函数，每一次返回一行文本数据，忽略行尾标记，所以返回的值可能是空值。一个行尾标记由可能存在的回车符和一个换行符组成，在正则表达式中为`\r?\n`，输入的 data 数据中的最后一行也会被返回，即使最后一行没有换行符

`func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanRunes 是 Scanner 的一个分割函数，每次返回一个 URF-8 字符作为 token。返回结果的 runes 的顺序和调用 range 进行访问一致，也就意味着错误的 UTF-8 编码会转换成`U+FFFD = "\xef\xbf\xbd".`。由于 Scan 接口，使得用户无法判断是编码替换还是编码错误。

`func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)`

ScanWords 是 Scanner 的一个分割函数，每次返回文本中一个由空白(定义见 unicode.IsSpace)分割的词，并且两边的空格会去掉。该函数不会返回空字符串。

## ReadWriter

```go
type ReadWriter struct {
	*Reader
	*Writer
}
```

ReadWriter 分别存储了一个 Reader 和 Writer 的指针，它实现了 io.ReadWriter 接口

`func NewReadWriter(r *Reader, w *Writer) *ReadWriter`

NewReadWriter 申请了一个读写器并且分发给了 r 和 w

## Reader

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader 实现了 io.Reader 缓存接口

`func NewReader(rd io.Reader) *Reader`

NewReader 返回一个缓存为默认大小的 Reader

`func NewReaderSize(rd io.Reader, size int) *Reader`

NewReaderSize 返回一个 Reader，该 Reader 的缓存大小至少有指定的 size。如果参数 rd 已经是一个 Reader 并 size 足够大，那么直接返回底层的 Reader 对象

```go
// Is it already a Reader?
b, ok := rd.(*Reader)
if ok && len(b.buf) >= size {
    return b
}
```

`func (b *Reader) Buffered() int`

Buffered 返回缓存中可以读取的字节数

`func (b *Reader) Discard(n int) (discarded int, err error)`

Discard 跳过之后的 n 个字节数据，返回跳过的字节数。如果 Discard 跳过的字节数少于 n 个字节，会返回错误。如果 0<= n<=b.Buffered()，Discard 可以在不从底层的 io.Reader 中读取数据就保证成功。

`func (b *Reader) Peek(n int) ([]byte, error)`

Peek 返回 Reader 下 n 个字节，但是不会移动读指针。返回的字节只有在下一次读取数据之前有效。如果 Peek 返回的字节数少于 n，那么会返回错误进行解释，如果 n>缓存字节数，那么返回错误 ErrBufferFull

`func (b *Reader) Read(p []byte) (n int, err error)`

Read 将数据读入到 p 中，该函数返回读入到 p 中的字节数，每次调用最多读取一个字节，因此 n 可能小于 len(p)，如果要准确的读道 len(p)个字节数，使用 io.ReadFull(b, p). 如果遇到 EOF 错误，返回的数量为 0，并且错误类型为 io.EOF(准确的来说，这不是一个错误，而是一个标志)

`func (b *Reader) ReadByte() (byte, error)`

ReadByte 每次读取一个字节，如果没有字节剩余，那么会返回一个错误

`func (b *Reader) ReadBytes(delim byte) ([]byte, error)`

ReadBytes 会从输入中不断读取数据，直到遇到第一个分隔符，然后返回包括分隔符在内的之前的数据。 如果 ReadBytes 在找到一个分隔符之前遇到了错误，会返回在该错误之前读取到的数据以及错误(通常为 io.EOF)。当前仅当数据不以 delim 分割的时候 ReadBytes 返回 err!=nil，对于简单的使用，使用 Scanner 会更加方便。

`func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)`

ReadLine 函数是一个低水平的行读取函数，绝大多数的调用者应该使用 ReadBytes("\n")或者 ReadString("\n")，或者使用 Scanner

ReadLine 会返回一整行，但是不会包含最后行末符号。如果该行超过了 buffer 的容量，那么会设置 isPrefix 为 true，并且返回行首内内容，剩余的部分会在之后的调用中返回，直到该行的内容全部返回，isPrefix 才会返回 false。返回的内容仅在下一次调用 ReadLine 之前有效。ReadLine 要么返回非空的一行数据，那么返回一个错误，绝对不会一起发生。

返回的文本内容不会包括行末("\r\n"和"\n")。如果输入结束的时候没有最后一行结束，不会返回任何只是和错误。在 ReadLine 之后调用 UnReadByte 将始终读取最后读取的字节，可能是行尾的字符，即便该字节不是 ReadLine 返回的一部分(可能是"\n")

`func (b *Reader) ReadRune() (r rune, size int, err error)`

ReadRune 读取一个 UTF-8 编码的字符，并且返回该字符以及它的字节数。 如果这个编码的 rune 无效，它消耗一个字节(读指针移动一个字节)，返回 unicode.ReplacementChar(U+FFFD)，size 值为 1

`func (b *Reader) ReadSlice(delim byte) (line []byte, err error)`

ReadSlice 读取数据直到遇到第一个 delim 字符，返回缓存中已经包含的内容，该值只在下一次读取数据之前有效。如果 ReadSlice 在遇到第一个 delim 字符之前发生了错误，它会返回缓冲中的内容以及 error 本身(通常为 io.EOF)。如果在读取到 delim 之前 buffer 已经满了，ReadSlice 会返回 ErrBufferFull 错误。因为返回的数据会被下一次的 IO 操作覆盖，所以大多数的用户应该调用 ReadBytes 或者 ReadString。ReadSlice 仅当返回的切片不以 delim 结尾时返回 err!=nil

`func (b *Reader) ReadString(delim byte) (string, error)`

ReadString 从输入中读取数据直到遇到第一个 delim 字符，返回包含已读取数据和 delim 的字符串。如果只是简单使用，Scanner 会更加方便

`func (b *Reader) Reset(r io.Reader)`

Reset 清空缓存中的数据，重置所有的状态，并且 r 作为读取数据的来源

`func (b *Reader) Size() int`

Size 返回底层缓存的大小

`func (b *Reader) UnreadByte() error`

UnreadByte 吐出最近一次读取数据的最后一个字节。 如果最近调用的函数不是读取操作，那么 UnreadByte 将会返回一个错误。注意，Peek 函数不算为一次读取操作

`func (b *Reader) UnreadRune() error`

UnreadRune 吐出上一次读取数据的最后一个 rune，如果租金一次调用函数不是 ReadRune 该函数会返回一个错误。在这一点上，UnreadRune 要比 UnreadByte 严格，UnreadByte 适用于任何读操作。

`func (b *Reader) WriteTo(w io.Writer) (n int64, err error)`

WriteTo 实现了接口 io.WriteTo。该函数可能多次调用底层 Reader 的读方法，如果底层的 reader 支持 WriteTo 方法，该函数的调用直接调用底层的 WriteTo 方法而不进行缓冲

## Scanner

```go
type Scanner struct {
	// contains filtered or unexported fields
}
```

Scanner 提供了一个简便的接口读取数据，比如说由换行符分割的文本文件。

`func NewScanner(r io.Reader) *Scanner`

NewScanner 创建一个从 r 中读取数据的 Scanner，默认的分割函数为 ScanLines

`func (s *Scanner) Buffer(buf []byte, max int)`

Buffer 设置搜索时初始的缓冲大小，缓冲大小的最小值可能在扫描中重新赋值。最大的 token 大小是 max 和 cap(buf)中的最大值，如果 max<=cap(buf),Scan 会使用这个缓冲并且不会再进行分配

在默认情况下，Scan 使用内置的缓冲，并且设置最大的 token 大小为 MaxScanTokenSize.

如果在扫描之后调用该函数会 panic

`func (s *Scanner) Bytes() []byte`

Bytes 会返回最近一次调用 Scan 生成的 token。 底层的数组可能会被下一次的 Scan 重写

`func (s *Scanner) Err() error`

Err 返回 Scanner 遇到的第一个非 EOF 的错误

`func (s *Scanner) Scan() bool`

Scann 会让 Scanner 的扫描位置到下一个 token 的位置，并且当前的 token 可以由 Bytes 或者 Text 方法获取。当扫描结束的时候(遇到错误或者到达结尾)，Scan 返回 false，Err 方法会返回扫描中遇到的错误，如果该错误为 io.EOF，Err 会返回 nil

`func (s *Scanner) Split(split SplitFunc)`

Split 方法设置 Scanner 的分割函数，必须在扫描之前调用该方法

Split 类型如下：

```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

package bufio 中已经含有几个预先定义的 Split 函数

`func (s *Scanner) Text() string`

Text 返回 Scan 方法最近一次产生的 token，会申请一个字符串保存该 token 并返回

## Writer

```go
type Writer struct {
	// contains filtered or unexported fields
}
```

Writer 实现了为 io.Writer 接口对象提供缓冲。如果 Writer 写的时候发生了错误，不会再读取任何数据并且任何写操作都会返回该错误。在所有的数据都写入之后，用户应该调用 Flush 函数来将数据都交给底层的 io.Writer

`func NewWriter(w io.Writer) *Writer`

NewWriter 返回一个缓冲为默认大小的 Writer

`func NewWriterSize(w io.Writer, size int) *Writer`

NewWriter 返回一个新的 Writer，且它的缓冲至少为指定的大小。如果 io.Writer 已经是一个 Writer(因为 Writer 其实是实现了 Write 方法的)，并且缓冲足够大，那么直接返回底层的 Writer 对象

`func (b *Writer) Available() int`

Available 返回缓冲中还有多少数据没有使用

`func (b *Writer) Buffered() int`

Buffered 返回目前缓冲中已经写入了多少的数据

`func (b *Writer) Flush() error`

Flush 将缓冲的数据写入到底层的 io.Writer 中

`func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)`

ReadFrom 实现了 io.ReadFrom. 如果底层的 writer 支持 ReadForm 方法，而且 b 中还没有缓冲的数据，那么会无缓冲的调用底层的 ReadFrom 方法

`func (b *Writer) Reset(w io.Writer)`

Reset 会清除掉所有没有进行 flush 的缓冲数据，删除所有的错误，然后将数据输出设为 w

`func (b *Writer) Size() int`

Size 返回底层缓冲的字节大小

`func (b *Writer) Write(p []byte) (nn int, err error)`

Write 将 p 中的内容写入到缓冲中，并且返回写入的字节数 n，如果 nn < len(p)，那么它也会返回一个错误解释原因。

`func (b *Writer) WriteByte(c byte) error`

WriteByte 写入单个字符

`func (b *Writer) WriteRune(r rune) (size int, err error)`

WriteRune 写入单个 Unicode 编码值，并且返回写入的字节数以及遇到的任何错误。

`func (b *Writer) WriteString(s string) (int, error)`

WriteString 写入一个 string，并且返回写入的字节数，如果小于 len(s)，则会返回错误解释
