package main

import("fmt")
//滑动窗口，左右指针维护这个窗口
func lengthOfLongestSubstring(s string) int {
	m:=make(map[byte]int)//开一个map,key为某个字符,value为其索引，方便后续移动左指针
    left,max:=0,0
	for right:=0;right<len(s);right++{//右指针遍历
		if i,exists:=m[s[right]];exists&&i>=left{//右指针移到了第i个元素,如果它已经存在于滑动窗口中（左右指针可以重合，因此是i>=left）
			left=i+1//就把左指针移动到这个元素上一次出现地方的下一位，保证窗口里没有重复字符
		}
		m[s[right]]=right//更新索引，对于刚刚右指针移到了第i个元素，其value索引就是right
		current:=right-left+1
		if current>=max{
			max=current
		}
	}
	return max
}

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
