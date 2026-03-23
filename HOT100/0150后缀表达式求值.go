package main

import("fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack:=[]int{}
    l:=len(tokens)
	for i:=0;i<l;i++{
		if tokens[i]=="+"{
			a:=stack[len(stack)-2]
			b:=stack[len(stack)-1]
			stack=stack[:len(stack)-2]
			stack=append(stack,a+b)
		}else if tokens[i]=="-"{
			a:=stack[len(stack)-2]
			b:=stack[len(stack)-1]
			stack=stack[:len(stack)-2]
			stack=append(stack,a-b)
		}else if tokens[i]=="*"{
			a:=stack[len(stack)-2]
			b:=stack[len(stack)-1]
			stack=stack[:len(stack)-2]
			stack=append(stack,a*b)
		}else if tokens[i]=="/"{
			a:=stack[len(stack)-2]
			b:=stack[len(stack)-1]
			stack=stack[:len(stack)-2]
			stack=append(stack,a/b)
		}else{
			num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
		}
	}
	return stack[0]
}

func main(){
	fmt.Println(evalRPN([]string{"2","1","+","3","*"}))
}