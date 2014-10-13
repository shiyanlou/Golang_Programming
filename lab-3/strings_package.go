package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "go_lang"
	fmt.Printf("T/F? Does the string \"%s\" have prefix \"%s\"? ", str, "go")
	fmt.Printf("%t\n", strings.HasPrefix(str, "go"))
	fmt.Printf("T/F? Does the string \"%s\" contains \"%s\"? ", str, "-")
	fmt.Printf("%t\n", strings.Contains(str, "-"))
	str_new := strings.Replace(str, "go", "python", 1)
	fmt.Printf("Origin string: \"%s\", after replace: \"%s\"\n", str, str_new)
	fmt.Printf("Number of 'n' in \"%s\" is: %d\n", str_new, strings.Count(str_new, "n"))

}
