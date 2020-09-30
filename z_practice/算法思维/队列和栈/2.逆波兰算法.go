package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for i:=0;i<len(tokens);i++{
		if !checkOpera(tokens[i]){
			if v,ok :=strconv.Atoi(tokens[i]); ok==nil{
				stack = append(stack, v)
			}else{
				return 0
			}
		}else{
			a := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			if tokens[i] == "+"{
				newv := a+b
				stack = append(stack, newv)
			}else if tokens[i] == "-"{
				newv := b-a
				stack = append(stack, newv)
			}else if tokens[i] == "*"{
				newv := a*b
				stack = append(stack, newv)
			}else if tokens[i] == "/"{
				newv := b/a
				stack = append(stack, newv)
			} else{
				return 0
			}
		}
	}
	return stack[0]
}

func checkOpera(s string) bool{
	opera := []string{"+", "-", "*", "/"}
	for _,v :=range opera{
		if v == s{
			return true
		}
	}
	return false
}

func main(){
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
}