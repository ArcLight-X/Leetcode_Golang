package main

import("fmt")

func twoSum(nums []int, target int) []int {
    //双重循环极易超时
	//对于每一个数Nums[i],找target-nums[i]，没找到就把当前数存入哈希表
	//用哈希表查找(go中即map)
	//index0 2 ;index 1 7 ....
	//map:{0,2},{1,7}这样存储
	//键为数值，值为索引的哈希表，容量为数组大小，自动扩容
	m:=make(map[int]int,len(nums))
	for index,value:=range nums{
		x:=target-value//相当于找哈希表里的键，避免第二次循环
		if j,ok:=m[x];ok{//ok是bool，表示是否存在这样的m[x],查找是哈希函数直接定位，避免第二次循环
			return []int{index,j}
		}
		m[value]=index//键值存储逻辑与常规相反，注意
	}
	return nil
}

func main(){
	fmt.Println(twoSum([]int{2,7,11,15},9))
}
