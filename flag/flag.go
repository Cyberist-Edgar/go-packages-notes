package main 

import(
	"fmt"
	"flag"
)

var(
	url string
	method string 
)

func init(){
	flag.StringVar(&url,"url", "", "指定访问的url")
	flag.StringVar(&method,"method", "GET", "指定请求方式")
}

func main(){
	// 解析所有的参数，一定要先调用，然后才能使用对应的值！！
	flag.Parse()
	fmt.Println("------------Visit------------------")
	// 该函数只会对命令行中定义好的参数进行调用，声明的不一定会调用
	flag.Visit(func(flag *flag.Flag){
		fmt.Println(flag)
	})
	fmt.Println("------------VisitAll---------------")

	// 该函数会对所有的命令行参数进行调用，即使在命令行中没有出现
	flag.VisitAll(func(flag *flag.Flag){
		fmt.Println(flag)
	})

	// 设置已经注册好的参数的名称
	flag.Set("url", "hello")
	fmt.Println(url)

	// 返回参数对应的Flag对象
	// 该对象在命令行参数中与method(下面例子中)绑定
	// 如果不存在则返回nil
	fmt.Println(flag.Lookup("method"))
	fmt.Println(flag.Lookup("not_exists"))
	
	// 返回没有被flag解析的参数，这些参数不能有命令行参数
	// 只是对应的值，且这些值只能够放在末尾
	fmt.Println(flag.Args())
	// 相当于返回len(flag.Args())
	fmt.Println(flag.NArg())
	// 相当于flag.Args()[0]，如果不存在，返回空字符串
	fmt.Println(flag.Arg(0))

	// 返回命令行参数中提到的已经注册的参数格式
	fmt.Println(flag.NFlag())
}
