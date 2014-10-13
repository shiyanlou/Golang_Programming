package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len and cap of array %v is: %d and %d\n", a, len(a), cap(a))
	fmt.Printf("item in array: %v is:", a)
	for _, value := range a {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s1 := a[3:6]
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s1, len(s1), cap(s1))
	fmt.Printf("item in slice: %v is:", s1)
	for _, value := range s1 {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s1[0] = 456
	fmt.Printf("item in array changed after changing slice: %v is:", s1)
	for _, value := range a {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s2 := make([]int, 10, 20)
	s2[4] = 5
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s2, len(s2), cap(s2))
	fmt.Printf("item in slice %v is:", s2)
	for _, value := range s2 {
		fmt.Printf("% d", value)
	}

	fmt.Println()
}
