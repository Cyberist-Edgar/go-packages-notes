# sync.Once 基本使用

Once 是一个尽管多次调用但是只执行一次的对象
## 方法使用
`func (o *Once) Do(f func())`


Do 方法当且仅当第一次被调用的时候才会执行函数 f，即使之后调用的 f 不同，也不会对其进行调用. Do可以用于只需要一次的初始化

## 示例代码

[once.go](once.go)
