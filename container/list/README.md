# container/list 基本使用
list 实现了一个双向链表

## type Element

```go
type Element struct {
    // 元素保管的值
    Value interface{}
    // 内含隐藏或非导出字段
}
```
Element类型代表是双向链表的一个元素。

`func (e *Element) Next() *Element`

Next返回链表的后一个元素或者nil。

`func (e *Element) Prev() *Element`

Prev返回链表的前一个元素或者nil。

## type List 
```go
type List struct {
    // 内含隐藏或非导出字段
}
```
`func (l *List) Len() int`

Len返回链表中元素的个数，复杂度O(1)。

`func (l *List) Front() *Element`

Front返回链表第一个元素或nil。

`func (l *List) Back() *Element`

Back返回链表最后一个元素或nil。

`func (l *List) PushFront(v interface{}) *Element`

PushBack将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。

`func (l *List) PushFrontList(other *List)`

PushFrontList创建链表other的拷贝，并将拷贝的最后一个位置连接到链表l的第一个位置,

`func (l *List) PushBack(v interface{}) *Element`

PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。

`func (l *List) PushBackList(other *List)`

PushBack创建链表other的拷贝，并将链表l的最后一个位置连接到拷贝的第一个位置。

`func (l *List) InsertBefore(v interface{}, mark *Element) *Element`

InsertBefore将一个值为v的新元素插入到mark前面，并返回生成的新元素。如果mark不是l的元素，l不会被修改。

`func (l *List) InsertAfter(v interface{}, mark *Element) *Element`

InsertAfter将一个值为v的新元素插入到mark后面，并返回新生成的元素。如果mark不是l的元素，l不会被修改。

`func (l *List) MoveToFront(e *Element)`

MoveToFront将元素e移动到链表的第一个位置，如果e不是l的元素，l不会被修改。

`func (l *List) MoveToBack(e *Element)`

MoveToBack将元素e移动到链表的最后一个位置，如果e不是l的元素，l不会被修改。

`func (l *List) MoveBefore(e, mark *Element)`

MoveBefore将元素e移动到mark的前面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

`func (l *List) MoveAfter(e, mark *Element)`

MoveAfter将元素e移动到mark的后面。如果e或mark不是l的元素，或者e==mark，l不会被修改。

`func (l *List) Remove(e *Element) interface{}`

Remove删除链表中的元素e，并返回e.Value。

## 示例代码
[list.go](list.go)