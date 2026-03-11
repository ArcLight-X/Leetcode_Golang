package main

import (
	"fmt"
)	
//双指针法，从两边扫描整个数组，每次移动矮的一边，因为面积受矮柱子制约
func maxArea(height []int) int {
    left,right:=0,len(height)-1
	max:=0
	for left<right{
	if height[left]<=height[right]{
		area:=(right-left)*min(height[right],height[left])
		left++
		if area>=max{
			max=area
		}

	}else{
		area:=(right-left)*min(height[left],height[right])
		right--
		if area>=max{
			max=area
		}
	}
}
	return max
}

func main(){
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
