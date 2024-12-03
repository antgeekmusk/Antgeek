package main

import "fmt"

// go 面相对象

// 定义一个猫对象

// go中 所有的命名首字母大写代表public 可以被其他包访问,小写代表private 只能在这个包使用
// 实现封装的方式就是将结构体的属性名都小写,然后提供set和get方法
type Cat struct {
	Name  string
	Age   int
	Color string
}

// 显示对象的信息 这里的this,谁调用就是指向的谁
func (this Cat) Show() {
	fmt.Println("Name : ", this.Name)
	fmt.Println("Age : ", this.Age)
	fmt.Println("Color : ", this.Color)
}

// 获取名称
func (this Cat) getName() string {
	return this.Name
}

// 设置名称(这里的设置名称需要写成 this *Cat,这样才可以引用传递,否则是值传递,对象的值不会改变)
func (this *Cat) setName(newName string) {
	this.Name = newName
}

func main() {
	// 创建对象
	var cat1 Cat
	cat1.Name = "Tom"
	cat1.Age = 3
	cat1.Color = "black"
	// 调用方法
	cat1.Show()
	c1Name := cat1.getName()
	println(c1Name)
	cat1.setName("Jack")
	println(cat1.getName())

	// 另外一种创建对象的方式
	c1 := Cat{
		Name:  "Jane",
		Age:   10,
		Color: "blue",
	}
	c1.Show()
}
