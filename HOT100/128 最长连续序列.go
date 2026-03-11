package main

import("fmt"

)

func longestConsecutive(nums []int) int {
    //构建哈希表，从起点查找
	//对于100，查找上一个99，不存在，100是起点
	//查找下一个101，不存在，序列长为1
	//对于4,查找上一个3，存在,4不是起点，直接跳过,等下3往后找一定能找到4
	m:=make(map[int]bool,len(nums))
	for _,num:=range nums{
		m[num]=true//把数组存入表中,key为nums中的数，value为bool,标记这个数在不在哈希表中
	}//存储完即{100,true},{4,true}....
	max:=0
	for num:=range m{
		if !m[num-1]{//上一个不存在,即是起点
			Currertnum:=num//记录当前索引
			len:=1
		for m[Currertnum+1]{//向后一个一个查找,若找到
			Currertnum++
			len++
		}
		if len>max{
			max=len
		}
		}
	}
	return max
}

func main(){
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
}
