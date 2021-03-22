package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int)bool{
	s :=strconv.Itoa(x)
	n := len(s)
	flag:=true
	fmt.Println("len",n)


	if n==1{
		return true //只有一个数的时候返回true
	} else {
		numSlice := strings.Split(s, "")
		if n%2==1{
			for i:=0;i<(n-1)/2;i++ {
				fmt.Println(numSlice[0])

				fmt.Println(numSlice[6])
				if numSlice[i]== numSlice[n-i-1]{
					flag=false
					println(flag)

				}else {
					flag= false
				}
			}

		}else {
			for i:=0;i<n/2;i++ {
				//fmt.Println(numSlice[i])

				if numSlice[i]== numSlice[n-i-1]{
					flag=true

				}else {
					flag= false
				}

			}

		}

	}
	return flag

}

func main() {

	var x int
	_, _ = fmt.Scanln(&x)
	//s = "1221"
	res:= isPalindrome(x)
	println(res)

}


