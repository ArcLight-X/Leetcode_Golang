package main

import("fmt"

)

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1:=headA
	p2:=headB
	//p1先在第一个链表走到底，然后从第二个链表开始走，p2相反，由对称性，p1=p2的点就是交点，要么就没交点
	for p1!=p2{
		if p1==nil{
			p1=headB
		}else{
			p1=p1.Next
		}
		if p2==nil{
			p2=headA
		}else{
			p2=p2.Next
		}
	}
	return p1
}
