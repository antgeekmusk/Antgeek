package main

import (
	"./lib1"
	"./lib2"
	"fmt"
)

func init() {
	fmt.Println("main init")
}
func main() {
	lib1.Test1()
	lib2.Test2()
}
