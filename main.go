package main

import (
	"fmt"
	"strconv"
)
// 第一个golang程序
// go run main.go com.go
func main() {
	printer("hello golang")
	mutirun(6)
	for i:=7; i<=14; i++ {
		mutirun(i)
	}

	printer("------------end---------------")
}
//顺序编程: 求一个数的阶乘
func mutirun(n int) {
	 str, total := "", 1
	for i:=1;i<=n;i++ {
		str += " "+strconv.Itoa(i)+" x"
		total *= i
	}
	res := fmt.Sprintf("求%02d的阶乘结果：%02d! = %s = %d", n, n, str[:len(str)-2], total)
	printer(res)
}