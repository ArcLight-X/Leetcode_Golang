package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//快慢指针：快指针先走n步，然后一起走，快指针到尾部时，慢指针就在倒数第n个
	dummy:=&ListNode{Next:head}//创建一个虚拟头节点，防止第n个正好是头节点
	slow:=dummy
	fast:=dummy
	for i:=0;i<=n;i++{
		fast=fast.Next
	}
	for fast!=nil{//确保fast走到Nil，slow才在倒数第n个
		fast=fast.Next
		slow=slow.Next
	}
	temp:=slow.Next
	slow.Next=temp.Next
	return dummy.Next
}
