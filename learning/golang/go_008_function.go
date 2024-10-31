package main

import "math"

// 函数
func main() {
	v1 := 1
	v2 := 2
	println("求最大", max(v1, v2))

	v3, v4 := swap(v1, v2)
	println("返回多个值", v3, v4)

	a := 10
	println("调用addOne之前的值", a)
	addOne(a)
	println("调用addOne之后的值", a)

	println("调用addTwo之前的值", a)
	addTwo(&a) // 将a的地址传递进去
	println("调用addTwo之后的值", a)

	funcParameterDemo()

	// 匿名函数测试
	// 返回一个函数
	nextNumber := anonymousFunctionDemo()
	println("匿名函数测试", nextNumber())
	println("匿名函数测试", nextNumber())
	println("匿名函数测试", nextNumber())
	// 新定义一个函数
	nextNumber2 := anonymousFunctionDemo()
	println("匿名函数测试2", nextNumber2())
	println("匿名函数测试2", nextNumber2())

	anonymousFunctionDemo2()

	methodDemo()

}

// 求最大值函数
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 返回多个值
func swap(a, b int) (int, int) {
	return b, a
}

// 值传递
func addOne(a int) int {
	return a + 1
}

// 引用传递
func addTwo(a *int) {
	// *a 就是取a 地址的值
	*a = *a + 2
}

// 函数作为另一个函数的实参
func funcParameterDemo() {
	// 定义一个函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 将函数作为参数传递
	println(getSquareRoot(16))

}

// 闭包
func anonymousFunctionDemo() func() int {
	// 定义一个变量
	i := 0
	return func() int {
		// 匿名函数内部可以直接使用外部函数的变量
		i += 1
		return i
	}
}

func anonymousFunctionDemo2() {
	// 加法函数
	add := func(num1, num2 int) int {
		return num1 + num2
	}

	// 乘法函数
	multiply := func(num1, num2 int) int {
		return num1 * num2
	}

	// 计算函数
	calculate := func(op func(int, int) int, x, y int) int {
		return op(x, y)
	}

	// 使用加法
	println(calculate(add, 1, 2))

	// 使用乘法
	println(calculate(multiply, 3, 4))
}

// 方法
func methodDemo() {

	var c Circle
	c.radius = 10
	println("圆的面积 = ", c.getArea())

}

// 定义一个结构体
type Circle struct {
	radius float64
}

// 该方法属于Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
