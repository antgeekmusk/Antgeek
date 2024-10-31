package main

import "fmt"

// 接口
func main() {
	// 通过接口实现多态
	var p Phone
	p = new(HwPhone)
	p.call()
	p = new(IPhone)
	p.call()
}

// 手机接口
type Phone interface {
	call()
}

// 华为手机
type HwPhone struct{}

// 苹果手机
type IPhone struct{}

// 华为手机实现接口Phone的方法
func (p HwPhone) call() {
	fmt.Println("三折叠华为手机")
}

// 苹果手机实现Phone接口的方法
func (p IPhone) call() {
	fmt.Println("苹果手机")
}
