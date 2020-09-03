# encoding/csv 基本使用

## csv文件介绍
csv文件是一种以`,`为字段分隔符的文件，常常第一行表示字段名，之后的行为对应的数据集合，每一行为每一独立的记录，最后的一行记录可以有换行符也可以没有

一个典型的csv文件如下:

```
id, name, age
1, Tom, 20
2, Jack, 30
```

## encoding/csv 包
在标准库encoding/csv中提供了对csv文件的基本操作，该package中声明了两个结构体对象`Reader`, `Writer`，分别支持csv文件的读取以及写入。

### 使用Reader进行读取操作
package 中提供了NewReader的方法，可以返回一个Reader对象，从源码中我们可以看出构造出来的Reader的分隔符为`,`，当然我们可以修改，但是在大多数情况下并没有必要
```go
// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Comma: ',',
		r:     bufio.NewReader(r),
	}
}
```

NewReader需要一个`io.Reader`类型的参数，常见的实现了该接口的类型有os.File、strings.Reader、bufio.Reader、bytes.Buffer、bytes.Reader，当然我们在该方法中使用最多的就是os.File了

创建好Reader对象之后，我们便可以调用其中的方法读取数据了，其中Read方法每次读取一条记录，ReadAll方法一次性读完所有记录，它们都有一个error类型的返回值，`如果为io.EOF表示文件读取完成，并不是发生错误`.

### 代码示例
[csv_reader.go](csv_reader.go)

### 使用Writer进行写操作
package中提供了NewWriter方法可以构造一个Writer对象，从源码中可以看出，默认以`,`为分隔符
```go
// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Comma: ',',
		w:     bufio.NewWriter(w),
	}
}
```
NewWriter需要接收一个`io.Writer`类型的参数，常见实现了该接口的类型有os.File、strings.Writer、bufio.Writer、bytes.Buffer、bytes.Writer,在这里我们使用os.File

然后调用Wirter对象的Write和WriteAll方法便可以写入数据，调用Flush方法可以将缓冲区的数据进入到文件中，调用Error方法可以返回Write和Flush执行时候遇到的问题

### 代码示例
[csv_writer](csv_writer.go)

[csv_reader](csv_reader.go)
