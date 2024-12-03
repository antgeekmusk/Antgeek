package main

import (
	"fmt"
	"time"
)

// use GODEBUG analyse
func main() {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("use GODEBUG analyse")
	}
}
