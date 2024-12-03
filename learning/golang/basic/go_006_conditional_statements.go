package main

import (
	"fmt"
	"time"
)

// 条件语句
func main() {
	ifElseStatement()
	switchStatement()
	selectStatement1()
	selectStatement2()
}

// if if else 语句
func ifElseStatement() {
	a := 10
	// 条件不需要带括号
	if a < 20 {
		println(a)
	}

	// 也可以带着括号
	if a < 20 {
		println(a)
	}

	// else if
	if a < 20 {
		println(a)
	} else if a < 15 {
		println(a)
	} else {
		println(a)
	}

}

// switch 语句
func switchStatement() {
	var grade = "B"
	var marks = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":

		println("优秀")
		println()
	case grade == "B":
		println("良好")
		println()
	case grade == "C":
		println("及格")
		println()
	case grade == "D":
		println("不及格")
		println()
	default:
		println("未知")
		println()
	}

	// Type Switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型 : %T\n", i)
	case int:
		fmt.Printf("x 的类型是int\n")
	case float64:
		fmt.Printf("x 的类型是float64\n")
	case func(int) float64:
		fmt.Printf("x 的类型是 func(int) float64\n")
	case bool, string:
		fmt.Printf("x 的类型是bool 或 string\n")
	default:
		println("未知型")
	}

	// fallthrough
	a := 10
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case a > 20:
		fmt.Println("3、case 条件语句为 false")
	case a > 30:
		fmt.Println("4、case 条件语句为 false")
	default:
		fmt.Println("6、默认 case")
	}

}

// select 语句
func selectStatement1() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 定义协程(goroutine) 延时5秒后给通道c1发送数据
	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	// // 定义协程(goroutine) 延时10秒后给通道c2发送数据
	go func() {
		time.Sleep(10 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// 这里没有default语句所以会等待其中一个通道返回数据
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 最终结果 先是c1 返回数据,后是c2 返回数据
}

func selectStatement2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "one"
		}
	}()

	go func() {
		for {
			c2 <- "tow"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("no data")
		}
		time.Sleep(1 * time.Second)
	}
}
