package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ori string = "123456"
	var i int
	var s string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	i, _ = strconv.Atoi(ori)
	fmt.Printf("The integer is: %d\n", i)
	i = i + 5
	s = strconv.Itoa(i)
	fmt.Printf("The new string is: %s\n", s)
}
