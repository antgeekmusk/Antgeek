package main

import "fmt"

// 结构体

type People struct {
	name string
	age  int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(People{"tom", 10})

	// 使用 key value 形式
	fmt.Println(People{age: 10, name: "jack"})

	// 字段可以为空
	fmt.Println(People{name: "ken"})

	// 访问结构体的成员
	var p1 People

	// 赋值
	p1.name = "antg"
	p1.age = 20

	// 打印
	fmt.Printf("姓名 : %s\n", p1.name)
	fmt.Printf("年龄 : %d\n", p1.age)

	printPeople(p1)
	changePeopleAge(&p1)
	printPeople(p1)
}

// 将结构体传入函数
func printPeople(people People) {
	fmt.Printf("姓名 : %s\n", people.name)
	fmt.Printf("年龄 : %d\n", people.age)
}

// 结构体指针
func changePeopleAge(people *People) {
	people.age += 1
}
