package main

import (
	"fmt"
	"math"
)

func ContvertIntToInt16(x int) int16 {
	if math.MinInt16 <= x && x <= math.MaxInt16 {
		return int16(x)
	}

	panic(fmt.Sprintf("%d is out of int16 range", x)) // 手动触发panic
}

func main() {
	i := ContvertIntToInt16(655567)
	fmt.Printf("%d", i)
}
