package main

import "fmt"

func main()  {
	//位运算的演示
	fmt.Println(2&3)

	fmt.Println(2|3)

	fmt.Println(2^3)

	fmt.Println(-2^2)

	a :=1 >> 2
	c :=1 << 2
	fmt.Println("a=",a ,"c=",c)

}
