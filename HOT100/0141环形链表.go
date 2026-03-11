package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    //快慢针：如果有环，快针一定追上慢针
	slow:=head
	fast:=head
	for fast!=nil&&fast.Next!=nil{//指针要走到头
		slow=slow.Next
		fast=fast.Next.Next
		//慢走一步，快走两步
		if slow==fast{
			return true
		}
	}
	return false
}
