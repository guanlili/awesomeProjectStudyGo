package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int)bool{
	s :=strconv.Itoa(x)
	n := len(s)
	flag:=false
	if n==1{
		return true //只有一个数的时候返回true
	} else {
		numSlice := strings.Split(s, "")
		for i:=0;i<n/2;i++ {
			if numSlice[i]== numSlice[n-i-1]{
				flag=true
			}else {
				flag= false
				break
				}
		}
	}
	return flag
}
func main() {
	var x int64
	_, _ = fmt.Scanln(&x)
	res:= isPalindrome(int(x))
	println(res)
}


