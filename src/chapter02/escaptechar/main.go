package main

import "fmt"

func main()  {
	//演示转义字符的使用 \t
	fmt.Println("tomjack")
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	fmt.Println("tom say\"i love you\"")
	//\r 回车，从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞厉害")

}