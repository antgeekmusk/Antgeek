package main

import "fmt"

// 定义一个学生父类
type Student struct {
	name string
	age  int
}

// 学生的方法
func (this Student) Print() {
	fmt.Println("name : ", this.name)
	fmt.Println("age : ", this.age)
}

// 定义一个小学生
type PrimaryStudent struct {
	Student // 继承父类
	hobby   string
}

// 重写父类方法
func (this PrimaryStudent) Print() {
	fmt.Println("name : ", this.name)
	fmt.Println("age : ", this.age)
	fmt.Println("hobby : ", this.hobby)
}

func main() {
	s := Student{
		name: "tom",
		age:  11,
	}
	ps := PrimaryStudent{
		Student: Student{
			name: "jack",
			age:  10,
		},
		hobby: "football",
	}
	s.Print()
	ps.Print()
}
