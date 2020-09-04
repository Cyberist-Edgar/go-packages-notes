# net/url 基本使用

package url 实现了 url 的解析以及查询字段的转义

## 方法使用

`func PathEscape(s string) string`

PathEscape 函数可以转义字符串使其能够安全的放置在一个 URL 路径中，该方法会根据需要将特殊字符替换成%XX

`func PathUnescape(s string) (string, error)`

PathUnescape 是 PathEscape 的反变换，与 QueryUnescape 类似，但是 PathUnescape 不会将`+`转义成`""`(空格)

`func QueryEscape(s string) string`

QueryEscape 函数可以转义字符使其能够放置在 URL 查询字段中

`func QueryUnescape(s string) (string, error)`

QueryUnescape 是 QueryEscape 的反变换

## URL 对象及其方法

URL 类型定义如下：

```go
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

一个 URL 对象表示一个解析过的 URL，可以说是 URI，一般形式为:

```
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```

`func Parse(rawurl string) (*URL, error)`

Parse 将一个 rawurl 解析成 URL 结构，其中 rawurl 可能是相对的(没有 host 的路径)，也有可能是绝对的路径。解析一个没有 scheme 但由 hostname 和 path 的路径是无效的，但是由于解析时的歧义性，不一定会返回错误

`func ParseRequestURI(rawurl string) (*URL, error)`

ParseRequestURI 将一个 rawurl 解析成 URL 结构，该方法假设 rawurl 是通过 HTTP 请求获取的，所有 rawurl 中应该是一个绝对的路径，并且没有锚点(#fragment)

`func (u *URL) EscapedFragment() string`

EscapedFragment 会返回经过转义之后的 u.Fragment. 一般而言，每一个锚点都有多种可能的转义形式。 如果 u.RawFragment 是 u.Fragment 一个合法的转义，那么直接返回 u.RawFragment，否则 EscapedFragment 自己计算转义后的值。String 方法使用 EscapedFragment 去构造它的结果。一般来说，我们应该调用 EscapedFragment 而不是直接读取 u.RawFragment

`func (u *URL) EscapedPath() string`

EscapedPath 会返回经过转义之后的 u.Path，与 EscapedFragment 类似

`func (u *URL) Hostname() string`

Hostname 返回 u.Host, 并且会去掉端口号。如果结果被方括号包围，就像 IPv6 一样，那么返回的结果会去掉方括号

`func (u *URL) IsAbs() bool`

IsAbs 判断该 URL 是否是绝对路径

`func (u *URL) MarshalBinary() (text []byte, err error)`

MarshalBinary 将 u.String()的结果转换成[]byte 类型

`func (u *URL) Parse(ref string) (*URL, error)`

Parse 方法类似上文提到的 Parse 方法

`func (u *URL) Port() string`

Port 方法返回 u.Host 中的端口号，不包括`:`，如果 u.Host 中包含一个无效的端口号，返回空字符串

`func (u *URL) Query() Values`

Query 方法解析了 RawQuery 并且返回相应的值，该方法会去掉不合法的键值对，如果要检查错误使用 ParseQuery

`func (u *URL) Redacted() string`

Redacted 方法类似 String 方法，但是会将其中的任何密码替换成"xxxxx"

`func (u *URL) RequestURI() string`

RequestURI 方法返回经过编码的`path?query` 或者 `opaque?query`, 可以用于 HTTP 请求

`func (u *URL) ResolveReference(ref *URL) *URL`

ResolveReference 方法会将一个 URL ref 路径引用以 URL u 路径为基础，计算出对应的绝对值路径

`func (u *URL) String() string`

String 方法将 URL 对象重新组成一个有效的 URL 字符串

`func (u *URL) UnmarshalBinary(text []byte) error`

UnmarshalBinary 方法将一个[]byte 切片类型解析成一个 URL 对象，并赋给 u

## Userinfo 对象及其方法

```go
type Userinfo struct {
	// contains filtered or unexported fields
}
```

Userinfo 是 URL 一个不可变的用户和密码的封装，现有的 Userinfo 值保证有用户名设置(RFC 2396 允许的，可能是空的)和可选的密码。

`func User(username string) *Userinfo`

以 username 创建一个 Userinfo 的指针对象，没有设置 password

`func UserPassword(username, password string) *Userinfo`

以 username 和 password 创建一个 Userinfo 的指针对象。这个功能只能对传统的网站使用， RFC 2396 建议不要使用这种方式进行解释 Userinfo，因为这会有安全风险

`func (u *Userinfo) Password() (string, bool)`

返回 Userinfo 的密码

`func (u *Userinfo) Username() string`

返回 Userinfo 的用户名

`func (u *Userinfo) String() string`

以`username[:password]`标准格式返回编码的用户信息

## Value 对象及其方法

```go
type Values map[string][]string
```

Values 将一个值映射到一个切片，它通常被使用为查询参数和表单数据，不想 http.Header，Values 对大小写敏感

`func ParseQuery(query string) (Values, error)`

ParseQuery 解析 URL 编码后的查询字符串并且返回一个 Values 类型的数据

`func (v Values) Add(key, value string)`

Add 函数将 value 添加到 key 中，如果对应的 key 存在，values 会追加到相应的值后面

`func (v Values) Del(key string)`

Del 函数删除与 key 相关的数据

`func (v Values) Encode() string`

Encode 函数以 key 排序，将值编码成 URLencoded("bar=baz&foo=quux")类的字符串

`func (v Values) Get(key string) string`

Get 方法返回所给 key 对应的第一个值，如果想要获取多个值，直接使用 map 类型的访问方式进行访问

`func (v Values) Set(key, value string)`

Set 方法会将 key 对应的值设置成 value，会覆盖之前的值


## 示例代码
[url.go](url.go)

[values.go](values.go)

[userinfo.go](userinfo.go)


