# errors 基本使用

package errors 中实现了对 error 类型数据的相关操作函数

## 概念

如果一个错误类型有一个 Unwrap 方法返回一个错误类型，则说该错误包装(wrap)另外一个错误。比如说 e.Unwrap()返回一个非 nil 的错误类型 w，那么我们说 e 包装(wrap)了 w

Unwrap 方法打开一个被包装的错误，如果它的参数类型含有 Unwrap 方法，它会调用该方法一次，否则它返回 nil

一个创建 wrapped errors 最简单的方法就是调用`fmt.Errorf`函数，并且使用`%w`占位符(%w 与%v 相同，只不过对于参数为 error 类型的数据，使用%w 会实现 Unwrap 方法)

```go
errors.Unwrap(fmt.Errorf("... %w ...", ..., err, ...))
// 这个方法返回的就是err本身
```

## 方法使用

`func As(err error, target interface{}) bool`

As 方法在 err 链中找到和 target 错误匹配的第一个错误，并且将该错误赋值给 target，并返回 true，否则返回 false

这个链由 err 本身和通过重复调用 Unwrap 获得的错误序列组成

As 方法会 panic 如果 target 不是一个指向实现了 error 的空指针类型或者接口

```go
// 下面两个方式都可以将perr定义为os.PathError
// 并赋值，但是使用第一种方式更好
var perr *os.PathError
if errors.As(err, &perr) {
	fmt.Println(perr.Path)
}

if perr, ok := err.(*os.PathError); ok {
	fmt.Println(perr.Path)
}

```

`func Is(err, target error) bool`

Is 方法判断是否存在在 err 链中的错误匹配 target

```go
// 第一种方法进行判断是否相等更好
if errors.Is(err, os.ErrExist)
if err == os.ErrExist
```

`func New(text string) error`

New 方法返回一个由所给的 text 格式化的错误类型。每一个 New 方法返回的错误都是不同的即使 text 一致

`func Unwrap(err error) error`

Unwrap 返回 err 调用 Unwrap 方法的结果值(如果该错误类型包含一个返回错误类型的 Unwrap 方法)，否则该方法返回 nil

## 示例代码

[errors.go](errors.go)
