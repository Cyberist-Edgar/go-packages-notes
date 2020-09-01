WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量，每一次执行Add都会增加线程组的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。

###### 方法介绍
1. `func (*WaitGroup) Add(delta int)`

Add方法向内部计数加上delta，`delta可以是负数`；如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。

2. `func (*WaitGroup) Done`

  Done方法减少WaitGroup计数器的值，应在线程的最后执行，Done的执行应标志着一个goroutine的结束

3. `func (*WaitGroup) Wait `

Wait方法阻塞直到WaitGroup计数器减为0。如果WaitGroup不为0，那么程序就会一直阻塞在Wait函数这里
