package main

import "fmt"

//给定一个字符串数组
//[“I”,“am”,“stupid”,“and”,“weak”]
//用 for 循环遍历该数组并修改为
//[“I”,“am”,“smart”,“and”,“strong”]

func main() {
	strArr := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("原先数组", strArr)
	strArrReplace := [5]string{"I", "am", "smart", "and", "strong"}
	for index, value := range strArr {
		// value只是副本不会改变
		value = strArrReplace[index]
		// 操作数组本身会改变
		strArr[index] = strArrReplace[index]
		fmt.Println(value)
	}
	fmt.Println("结果", strArr)
}
