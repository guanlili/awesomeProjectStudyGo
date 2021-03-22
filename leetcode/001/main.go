package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	n :=len(nums)
	//ans := make([]int, 2)

	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if nums[i]+nums[j]==target {
				//fmt.Println("i=,j=",i,j)
				//ans[0]=i
				//ans[1]=j
				return []int{i,j}

			}
		}
	}

	return nil
}

func twoSumhash(nums []int, target int) []int {

	hashTable:=map[int]int{}
	for i,x:=range nums{
		if p,ok:=hashTable[target-x];ok {
			return []int{p,i}
		}
		hashTable[x]=i
	}

	return nil
}
func main() {
	//fmt.Println("输入：")
	target := 9
	var nums =[]int{2,7,11,15}
	ret:= twoSumhash(nums,target)
	fmt.Printf( "输出 : %v\n", ret)

}