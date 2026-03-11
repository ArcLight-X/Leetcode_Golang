package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
    var x int
	for _,c:=range nums{
		x=x^c//直接异或即可
	}
	return x
}

func main(){
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
}
