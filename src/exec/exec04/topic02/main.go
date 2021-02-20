package main

import "fmt"

//对学生成绩大于60分的，输出合格，低于60分的，输出不合格（注，输入的成绩不能大于100）
func main()  {
	var score float64

	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	switch int(score /60 ){
	case 1:
		fmt.Println("合格")
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入错误")




	}


}
