package main

import "fmt"

// 封装例子

type Dog struct {
	name string // 小写首字母 权限private
}

func (this Dog) ToString() {
	fmt.Println("name :", this.name)
}

// get方法 方法首字母大写 public
func (this Dog) GetName() string {
	return this.name
}

// set方法
func (this *Dog) SetName(newName string) {
	this.name = newName
}

func main() {
	d := Dog{name: "jack"}
	d.ToString()
}
