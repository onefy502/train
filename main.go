package main

import (
	"fmt"
	"strconv"
)
// 第一个golang程序
// go run main.go com.go
func main() {
	printer("hello golang")
	Mutirun(6)
	for i:=7; i<=14; i++ {
		Mutirun(i)
	}
	//---
	bubble := []int{5,1,9,7,2,4,2,0,6,9,3}
	BubbleSort(bubble)
	printer(fmt.Sprint(bubble))
	Mutiply()
	printer("------------end---------------")
}
//顺序编程: 求一个数的阶乘
func Mutirun(n int) {
	 str, total := "", 1
	for i:=1;i<=n;i++ {
		str += " "+strconv.Itoa(i)+" x"
		total *= i
	}
	res := fmt.Sprintf("求%02d的阶乘结果：%02d! = %s = %d", n, n, str[:len(str)-2], total)
	printer(res)
}

//冒泡算法
func BubbleSort(arr []int) {
	n := len(arr)
	for i:=0;i<n;i++ {
		for j:=i+1; j<n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//九九乘法表
func Mutiply() {
	 res := ""
	for i:=1;i<=9;i++ {
		for j:=1;j<=i;j++ {
			res += fmt.Sprintf("%2d x %2d = %2d  ", j, i, i*j)
		}
		res += "\n"
	}
	printer(res)
}
