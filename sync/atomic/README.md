## sync/atomic
首先，我们来理解一下什么是`原子操作`，所谓原子操作就是指该操作是不可分割的，在执行完毕之前不会被其他的任务或者事件中断，该操作的内部不能被高层的操作发现并分割执行，处理器会首先保证我们这些操作内存的原子性，也就是说当向该内存地址写入或者读取数据的时候，其他的操作无法获取该内存的地址，也就是保证了该内存所保存数据的正确性。<br />
<br />在sync/atomic中提供了许多原子操作支持，主要有五大类：

### Load
load 一类的方法接收一个地址，并返回该地址中的数据<br />该类方法主要负责从相应的内存地址中获取对应的值
```go
return *addr
```

### Store
Store一类的方法接收两个参数，一个是保存数据的地址，另外一个是要保存的值<br />该类主要负责将对应的值保存在相应的内存地址中
```go
*addr = newValue
```

### Add
Add一类的方法接收两个参数，一个是保存数据的地址，一个是需要加上的数据，而后返回一个处理后的值<br />该类方法可以理解是Load和Store的结合，也就是先Load然后Add
```go
*addr += delta
return *addr
```

### Swap
Swap一类的方法接收两个参数，一个是保存数据的地址，一个是新的数据，函数会返回该地址原本保存的数据<br />该类方法可以理解为先Load，在Store新值，然后返回旧值
```go
oldValue = *addr
*addr = newValue
return oldValue
```

### CompareAndSwap
CompareAndSwap一类的方法接收三个参数，保存数据的地址，旧数据，新数值，函数会返回一个bool类型的结果<br />该类方法可以这样理解：先比较旧数据和地址中保存数据的值，如果相同的话，执行Swap，把新的数值保存在地址中，返回true，如果不同，那么直接返回false
```go
if *addr == old {
 *addr = new
 return true
}
return false
```

<br />在sync/atomic中还定义了一个Value对象，该对象实现了Load和Store两个方法，其中Load方法返回一个interface{}结果，而Store方法接收一个interface{}参数，其使用方法也就类似上述的Load和Store类的方法

### 示例代码
[atomic.go](atomic.go) 

[atomic_value.go](atomic_value.go)
