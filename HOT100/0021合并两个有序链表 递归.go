package main

import("fmt"

)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//递归法！非迭代法
	//一个链表为空就不用合并
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val { //比较两个链表头节点
		list1.Next = mergeTwoLists(list1.Next, list2) //把较小的那个头节点排在最前，剩下还是合并有序链表
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}


