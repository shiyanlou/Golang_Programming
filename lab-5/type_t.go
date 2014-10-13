package main

import (
	"fmt"
)

func main() {
	x := uint16(65000)
	y := int16(x) // 将 x转换为int16类型
	fmt.Printf("type and value of x is: %T and %d\n", x, x)
	fmt.Printf("type and value of y is: %T and %d\n", y, y)

	var i interface{} = 99 // 创建一个interface{}类型，其值为99
	var s interface{} = []string{"left", "right"}
	j := i.(int) // 我们假设i是兼容int类型，并使用类型断言将其转换为int类型
	fmt.Printf("type and value of j is: %T and %d\n", j, j)

	if s, ok := s.([]string); ok { // 创建了影子变量，if的作用域中覆盖了外部的变量s
		fmt.Printf("%T -> %q\n", s, s)
	}
}
