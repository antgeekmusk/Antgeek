package main

import "fmt"

// 多态例子

type Animal interface {
	Sleep()           // 动物睡觉
	GetColor() string // 获取动物颜色
	GetType() string  // 获取动物类型
}

// 定义一个结构体 pig
type Pig struct {
	color string
}

// 实现了接口的所有方法即实现了接口(必须全部实现)

func (this *Pig) Sleep() {
	fmt.Println("pig is sleep")
}
func (this *Pig) GetColor() string {
	return this.color
}
func (this *Pig) GetType() string {
	return "pig"
}

// 定义一个结构体 cow
type Cow struct {
	color string
}

func (this *Cow) Sleep() {
	fmt.Println("cow is sleep")
}
func (this *Cow) GetColor() string {
	return this.color
}
func (this *Cow) GetType() string {
	return "cow"
}

// 将对象传入实现不同类型对象调用不同的方法
func animalSleep(animal Animal) {
	animal.Sleep()
	fmt.Println("color : ", animal.GetColor())
	fmt.Println("type : ", animal.GetType())
}

func main() {
	var animal Animal
	animal = &Pig{color: "pink"} // 多态
	animal.Sleep()
	animalSleep(animal)
	animal = &Cow{color: "black"} // 多态
	animal.Sleep()
	animalSleep(animal)

}
