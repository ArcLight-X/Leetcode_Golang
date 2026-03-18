package main

import("fmt"
)

type MinStack struct {//构造两个栈
    elements []int//正常
	min []int//最小
}


func Constructor() MinStack {
    return MinStack{
		elements:[]int{},
		min:[]int{},
	}
}

//带*才是在原栈上修改
func (this *MinStack) Push(val int)  {
    this.elements=append(this.elements,val)
	//如果最小栈是空的或者入栈的值比当前最小的还小,入最小栈
	//确保最小栈的栈顶是最小值
	if len(this.min)==0||val<=this.min[len(this.min)-1]{
		this.min=append(this.min,val)
	}
}


func (this *MinStack) Pop()  {
	//如果栈顶是最小值，最小栈的也要弹出
	if this.elements[len(this.elements)-1]==this.min[len(this.min)-1]{
		this.min=this.min[:len(this.min)-1]
	}
    this.elements=this.elements[:len(this.elements)-1]	
}


func (this *MinStack) Top() int {
    return this.elements[len(this.elements)-1]
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}