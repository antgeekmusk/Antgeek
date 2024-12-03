package lib1

import "fmt"
import _ "../lib2" // 匿名import 使用这种语法可以导入但是不使用该包
func Test1() {
	fmt.Println("lib1 Test1")
}
func init() {
	fmt.Println("lib1 init")
}
