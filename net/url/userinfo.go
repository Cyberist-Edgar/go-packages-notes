package main

import (
	"fmt"
	"net/url"
)

func main(){
	// 创建一个Userinfo对象
	user := url.UserPassword("username", "password")

	fmt.Println("username: ", user.Username())
	fmt.Println("password: ", user.Password())
	fmt.Println(user.String())

}