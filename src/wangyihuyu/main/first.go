// 本题为考试多行输入输出规范示例，无需提交，不计分。
package main

import (
"fmt"
)
func main() {
	k:=0
	b:=0
	a:=0
	var numA [10]int
	var numB [10]int


	fmt.Scan(&a,&b,&k)
	for i := 0; i < 2; i++ {
		for j := 0; j < a; j++ {
			//x:=0
			fmt.Scan(&numA[j])
			//ans = ans + x
		}
	}

	fmt.Println(numA, numB)
	ans:=a+b+k
	fmt.Printf("%d\n",ans)
}