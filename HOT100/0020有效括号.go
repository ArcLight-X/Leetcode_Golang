package main

import("fmt")
/*Go 用切片模拟栈：
栈顶 = 切片最后一个元素
入栈 = append
出栈 = 切片截取 [:len(stack)-1]
判空 = len(stack) == 0
取栈顶 = stack[len(stack)-1]*/
func isValid(s string) bool {
    stack:=[]byte{}//字符栈
	for _,c:=range s{
		if c=='('||c=='['||c=='{'{
			stack=append(stack,byte(c))//入栈
		}else{
			if(len(stack)==0){//没有左括号入栈，肯定无效
				return false
			}
			top:=stack[len(stack)-1]
			if c==')'&&top=='('||c==']'&&top=='['||c=='}'&&top=='{'{//匹配上了
				stack=stack[:len(stack)-1]//栈顶元素弹出
			}else{//有一个没匹配上，就无效
				return false
			}
		}
	}
	return len(stack)==0//循环结束且栈空了才有效
}

func main(){
	fmt.Println(isValid("()[]{}"))
}

