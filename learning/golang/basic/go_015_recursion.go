package main

import "fmt"

// 递归
func main() {
	println("阶乘函数", factorial(50))
	var i uint64
	for i = 0; i < 10; i++ {
		fmt.Printf("斐波那契数列 : %d %d\n", i, fibonacci(i))

	}

	println("平方根", sqrt(24))

}

// 阶乘函数
func factorial(n uint64) (res uint64) {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// 求平方根
func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 1.0, 1e-9)
}

func sqrtRecursive(x, guess, preGuess, epilon float64) float64 {
	if diff := guess*guess - x; diff < epilon && -diff < epilon {
		return guess
	}
	// 牛顿法求平方根公式
	newGuess := (guess + x/guess) / 2
	if newGuess == preGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epilon)
}
