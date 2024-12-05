package main

import (
	"fmt"
	"goinit/lib1"
	"goinit/lib2"
)

func init() {
	fmt.Println("main init")
}
func main() {
	lib1.Test1()
	lib2.Test2()
}
