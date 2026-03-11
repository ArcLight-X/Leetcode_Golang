package main

import("fmt")

func threeSum(nums []int) [][]int {
	result := [][]int{}
    //先排序，之后遇到重复数字可以直接跳过，方便去重
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums)-1;j++{
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
			}
		}
	}
	//双指针法：先固定一个nums[k]，在nums[k+1]到nums[len-1]区间里找nums[left]和[right]
	//其中nums[left]+nums[right]=-nums[i]
	for k:=0;k<len(nums)-2;k++{//k+1，len-1区间至少还要有两个数才能双指针
		if k > 0 && nums[k]==nums[k-1] {//去重：如果当前数字与前一个相同，跳过，避免重复解
            continue
        }
		if nums[k]>0{//第一个数大于0了，后面都是正数不可能和是0
			break
		}
		left,right:=k+1,len(nums)-1
		for left<right{
		if nums[left]+nums[right]+nums[k]==0{//正好，双指针同时收缩
			result=append(result,[]int{nums[k], nums[left], nums[right]})
			for left<right&&nums[left]==nums[left+1] {//去重：跳过左侧重复的数字
                    left++
            } 
            for left<right&&nums[right]==nums[right-1] {//去重：跳过右侧重复的数字
                    right--
            }
			left++
			right--
		}else if nums[left]+nums[right]+nums[k]>0{//大了
			right--
		}else if nums[left]+nums[right]+nums[k]<0{//小了
			left++
		}
	}
	}
	return result
}
func main(){
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}
