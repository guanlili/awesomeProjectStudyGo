package main

import "fmt"

func partition(s string) [][]string {
	var stack []int
	stack.append(s[1])
	stack.pop()


	return nil
}

func main() {
	//fmt.Println("输入：")
	s :="aab"
	ret:= partition(s)
	fmt.Printf( "输出 : %v\n", ret)

}