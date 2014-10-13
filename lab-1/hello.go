package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "World"
	if len(os.Args) > 1 { /* os.Args[0]是“hello" */
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
