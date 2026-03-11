package main

import("fmt")

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    if head==nil||head.Next==nil{
		return nil
	}
	slow:=head
	fast:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next//同环形链表1，先找有没有环
		if slow==fast{//相遇则有环,此时 头到入环点 = 相遇点到入环点
			slow=head//慢针重置到表头
			for slow!=fast{
				slow=slow.Next
				fast=fast.Next//链表头到入环点的距离=第一次相遇点到入环点的距离
			}
			return slow
		}
	}
	return nil

}
