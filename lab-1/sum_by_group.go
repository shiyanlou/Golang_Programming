package main

import "fmt"
import "time"

func sum(a []int, result chan int) {
	time.Sleep(time.Second * 300)
	sum := 0
	for _, v := range a {
		sum += v
	}
	result <- sum
}

func main() {
	a := []int{2, 3, 5, 6, 10, -5, 1, 0}
	result := make(chan int)
	go sum(a[:len(a)/2], result)
	go sum(a[len(a)/2:], result)
	x, y := <-result, <-result

	fmt.Println(x, y, x+y)
}
