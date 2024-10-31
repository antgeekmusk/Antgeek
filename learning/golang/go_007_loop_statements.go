package main

import (
	"fmt"
	"time"
)

// 循环
func main() {
	//demo1()
	//demo2()
	//demo3()
	//demo4()
	//demo5()
	//demo6()
	//demo7()
	//demo8()
	//demo9()
	//demo10()
	demo11()
}

// 实现 1到10的数字之和
func demo1() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	println(sum)
}

// sum=1 自累加10次
func demo2() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	println(sum)

	sum = 1
	for sum <= 10 {
		sum += sum
	}
	println(sum)
}

// 无限循环
func demo3() {
	sum := 0
	for {
		sum++
		time.Sleep(1 * time.Second)
		println("无限循环", sum)
	}
	// 这行不会执行
	println(sum)
}

// for-each range 循环
func demo4() {
	// 定义字符串切片(长度不确定的都是切片)
	strs := []string{"hello", "world"}
	for index, str := range strs {
		fmt.Println(index, str)
	}

	// 定义数字类型数组
	nums := [5]int{1, 2, 3}
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// 定义map
	map1 := make(map[int]string)
	map1[0] = "tom"
	map1[1] = "jack"
	map1[6] = "antgeek"

	// 遍历map
	for k, v := range map1 {
		println(k, v)
	}

	// 只遍历k
	for k := range map1 {
		println(k)
	}

	// 遍历 value
	for _, v := range map1 {
		println(v)
	}

}

// 多重循环中 break 跳出标记循环
func demo5() {
	// 不使用标记跳出内层循环
	for i := 1; i <= 3; i++ {
		println(i)
		for i1 := 11; i1 <= 13; i1++ {
			println(i1)
			break
		}
	}

	// 使用标记直接跳出外层循环
outerLoop:
	for i := 1; i <= 3; i++ {
		println(i)
		for i1 := 11; i1 <= 13; i1++ {
			println(i1)
			break outerLoop
		}
	}
}

// switch 语句使用 break,默认每个case后面都有一个break,不需要自己显示的添加
func demo6() {
	num := 1
	switch num {
	case 5:
		println(5)
	case 3:
		println(3)
	case 1:
		println(1)
		// fallthrough 不可以写在break上面
		break
		fallthrough // 可以写fallthrough,但是下面的case也不会执行
	case 2:
		println(2)
	default:
		println(999)
	}
}

// 在select语句中使用 break

func demo7() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "world"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
		break
		// break 后面代码不执行
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

// 对于 for select ,break 仅可以跳出select,不可以跳出 for
func demo8() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "hello"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "world"
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			break // 这里的break不会跳出for循环,待会儿还会执行一遍select,如果想要跳出for循环可以使用return 或者break label 或者goto
			// break 后面代码不执行
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
		time.Sleep(1 * time.Second)
		println("for循环")
	}
}

// 使用 break label 形式跳出for select循环
func demo9() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "hello"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "world"
		}
	}()

forCode:
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			break forCode
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
		time.Sleep(1 * time.Second)
		println("for循环")
	}
	println("循环执行完毕")
}

// continue label
func demo10() {
	// 不使用label
	for i := 1; i <= 3; i++ {
		println("不使用label 的continue", i)
		for j := 11; j <= 13; j++ {
			if j == 12 {
				continue
			}
			println("不使用label 的continue", j)
		}
	}

	// 使用label
outerFor:
	for i := 1; i <= 3; i++ {
		println("使用label 的continue", i)
		for j := 11; j <= 13; j++ {
			if j == 12 {
				continue outerFor
			}
			println("使用label 的continue", j)
		}
	}
}

// goto demo 使用goto来实现一个continue的功能
func demo11() {
	a := 10
loop:
	for a <= 20 {
		if a == 15 {
			a++
			goto loop
		}
		println(a)
		a++
	}
}
