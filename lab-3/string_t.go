package main

import (
	"fmt"
)

func main() {
	t0 := "\u6B22\u8FCE\u6765\u5230" // t0内容：欢迎来到
	t1 := "\u5B9E\u9A8C\u697C"       // t1内容：实验楼
	t2 := t0 + t1
	for index, char := range t2 {
		fmt.Printf("%-2d    %U      '%c'    %X      %d\n",
			index, char, char, []byte(string(char)), len([]byte(string(char))))
	}
	fmt.Printf("length of t0: %d, t1: %d, t2: %d\n", len(t0), len(t1), len(t2))
	fmt.Printf("content of t2[0:2] is: %X\n", t2[0:2])
}
