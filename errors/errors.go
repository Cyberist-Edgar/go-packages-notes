package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("test errors")
	wraped := fmt.Errorf("... %w ...", err)
	fmt.Println("wrap: ", wraped)
	fmt.Println("Unwrap: ", errors.Unwrap(wraped))
	fmt.Println(err == wraped) // false
	// 注意wraped应该放在前面，因为会调用参数的Unwrap方法
	fmt.Println(errors.Is(wraped, err)) // true

	var e error
	// 在这里可以使用类型断言
	if errors.As(wraped, &e) {
		fmt.Println("e: ", e)
	}

}
