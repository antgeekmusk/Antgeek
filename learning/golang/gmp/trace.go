package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// 使用 go tool trace file.out 来分析gmp
func main() {
	// create file
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	// close file
	defer f.Close()

	// start trace
	trace.Start(f)

	// your code
	fmt.Println("hello trace")

	// stop trace
	trace.Stop()

}
