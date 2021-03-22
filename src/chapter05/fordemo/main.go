package main

import (
	"fmt"
)

func main()  {
	//输出10句你好，世界
	for i := 1; i <= 10; i++{
		fmt.Println("hello wrold",i)
	}

	//for循环的第二种写法
	j := 1
	for j<=10 {
		fmt.Println("nihao",j)
		j++

	}
	//for循环的第三种写法,这种写法通常配合break
	k :=1
	for ; ; {
		if k <=10{
			fmt.Println("ok",k)
		}else {
			break
		}
		k++
	}

	//字符串便利方式1-传统方式
	var str string = "hello,world!"
	for i:=0;i<len(str);i++ {
		fmt.Printf("%c\n",str[i])//使用下标
	}
	fmt.Println()
	//字符串便利方式2-for-range
	str = "abc ok"
	for index,val:=range str {
		fmt.Printf("index=%d,val=%c\n",index,val)
	}

}
