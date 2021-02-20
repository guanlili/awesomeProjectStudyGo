package main

import "fmt"

//使用swtich把小写类型的char型转为大写（键盘输入）。只转换a,b,c,d,e其他的输出other
func main()  {

	var char byte
	fmt.Println("请输入a,b,c,d,e")

	fmt.Scanf("%c" ,&char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}

}
