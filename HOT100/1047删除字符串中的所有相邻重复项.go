package main

import("fmt"
)

func empty(stack []byte) bool{
	return len(stack)==0//判空
}
func removeDuplicates(s string) string {
	stack:=[]byte{}//字符栈
	for i:=0;i<len(s);i++{
		if empty(stack)||s[i]!=stack[len(stack)-1]{//栈为空或者新入栈元素和栈顶不同
			stack=append(stack,byte(s[i]))//入栈
		}else{
			stack=stack[:len(stack)-1]//否则不入栈，并且栈顶一起弹出，实现删除相邻字符
		}
	}
	return string(stack)

}

func main(){
	fmt.Println(removeDuplicates("abbaca"))
}