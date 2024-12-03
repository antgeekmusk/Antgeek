package main

import (
	"errors"
	"fmt"
)

// 错误处理
func main() {
	res, err := errorSqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	errorDemo1()

}

// 使用errors.New 来抛出异常
func errorSqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("error : 参数不可以小于0")
	}
	// 实现
	return 0, nil
}

// 自定义实现error接口

// 定义一个 error结构体
type DivideError struct {
	dividee int
	divider int
}

// 实现error接口
func (de *DivideError) Error() string {
	strFormat := `
error :
Cannot proceed, the divider is zero.
dividee: %d
divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 除法运算函数
func Divide(dee int, der int) (res int, errorMsg string) {
	if der == 0 {
		dData := DivideError{
			dividee: dee,
			divider: der,
		}
		errorMsg = dData.Error()
		return
	} else {
		return dee / der, ""
	}
}

func errorDemo1() {
	// 正常情况
	if res, errMsg := Divide(100, 10); errMsg == "" {
		fmt.Println("100/10=", res)
	}
	if _, errMsg := Divide(100, 0); errMsg != "" {
		fmt.Println("100/0", errMsg)
	}
}
