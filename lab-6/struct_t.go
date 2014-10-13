package main

import "fmt"

type Count int // 创建自定义类型 Count

func (count *Count) Increment()  { *count++ } // Count类型的方法
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

type Part struct { // 基于结构体创建自定义类型 Part
	stat  string
	Count // 匿名字段
}

func (part Part) IsZero() bool { // 覆盖了匿名字段Count的IsZero()方法
	return part.Count.IsZero() && part.stat == "" // 调用了匿名字段的方法
}

func (part Part) String() string { // 定义String()方法，自定义了格式化指令%v的输出
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
	var i Count = -1
	fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zero t/f? : %t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")
	part := Part{"232", 0}
	fmt.Printf("Part: %v\n", part)
	fmt.Printf("Part is zero t/f? : %t\n", part.IsZero())
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero()) // 尽管覆盖了匿名字段的方法，单还是可以访问

}
