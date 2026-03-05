package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//不用反转链表，加法本来就从低位开始加
    ans:=&ListNode{}//listnode是结构体，&取地址。创建一个空的、值为0的链表头节点作为结果，并且拿到它的指针（地址）。
	curr:=ans
	carry:=0
	for l1!=nil||l2!=nil||carry!=0{//计算每一位的和
		sum:=carry
		if l1!=nil{
			sum=sum+l1.Val
			l1=l1.Next
		} 
		if l2!=nil{
			sum=sum+l2.Val
			l2=l2.Next
		}
		carry=sum/10//进位
		curr.Next=&ListNode{Val:sum%10}//本位
		curr=curr.Next
	}
	return ans.Next//头节点本身数据不需要
}