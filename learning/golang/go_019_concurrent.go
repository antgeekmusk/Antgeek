package main

import (
	"fmt"
	"time"
)

// 并发
func main() {
	//concurrentDemo1()
	//channelDemo1()
	//channelDemo2()
	selectDemo()

}

// goroutine 使用方法
func concurrentDemo1() {
	// 开启一个 goroutine
	go say("hello")
	say("world")
}

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 通道
func channelDemo1() {
	// 记录当前的时间
	start := time.Now()
	// 定义要求和的数组
	s := []int{1, 2, 3, 4, 5, 6}

	// 定义一个通道用于接收返回结果
	c := make(chan int)

	// 开启两个goroutine来并行计算这个任务

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// 从通道接收数据
	x, y := <-c, <-c
	fmt.Println("数据处理完成", x, y, x+y)

	// 打印花费时间
	fmt.Printf("处理时长 : %s\n", time.Since(start))

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		// 模拟任务延迟
		time.Sleep(1000 * time.Millisecond)
		sum += v
	}
	// 将sum 通过通道发送出去
	c <- sum
}

// go 遍历通道与关闭通道
func channelDemo2() {
	// 带有10个缓存单位的通道
	c := make(chan int, 10)
	go fibonacciChannel(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

// 通过channel 来实现斐波那契数列
func fibonacciChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// 发送完数据关闭
	close(c)
}

// select demo
func selectDemo() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
