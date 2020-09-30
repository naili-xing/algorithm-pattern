package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
394. 字符串解码
	给定一个经过编码的字符串，返回它解码后的字符串。
	编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
	你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
	此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
	示例 1：
		输入：s = "3[a]2[bc]"
		输出："aaabcbc"
	示例 2：
		输入：s = "3[a2[c]]"
		输出："accaccacc"
	示例 3：
		输入：s = "2[abc]3[cd]ef"
		输出："abcabccdcdcdef"
	示例 4：
		输入：s = "abc3[cd]xyz"
		输出："abccdcdcdxyz"
 */

func decodeString(s string) string {

	var stack []string

	for i:=0;i<len(s);i++{
		switch string(s[i]) {
		case "]":
			var temp []string

			//一旦遇到 "]" 就回退，开始找完整字符串，此时temp里的是倒序的
			for len(stack)>0 && stack[len(stack)-1]!="["{
				temp = append(temp, stack[len(stack)-1])
				stack = stack[0:len(stack)-1]
			}
			// 正序temp
			tempal := ""

			for i:=len(temp)-1;i>=0;i--{
				tempal += temp[i]
			}

			//找到后，把'['弹出
			stack = stack[0:len(stack)-1]

			//找这个字符串所对应的完整数字
			var numindex int = 1

			for len(stack)-numindex>=0 && isDigital(stack[len(stack)-numindex]) {
					numindex += 1
			}
			numsstack := stack[len(stack)-numindex+1: ]
			stack = stack[:len(stack)-numindex+1]
			nums := ""

			for i:=0;i<len(numsstack);i++{
				nums+=numsstack[i]
			}
			num,_ := strconv.Atoi(nums)

			//加入栈
			for i:=0;i<num;i++{
				stack = append(stack, tempal)
			}

		default:
			stack = append(stack, string(s[i]))

		}
	}

	return strings.Join(stack,"")
}



func isDigital(s string) bool{
	res := []string{"1", "2","3","4","5","6","7","8","9", "0"}
	for _,v := range res{
		if v == s{
			return true
		}
	}
	return false
}


func main(){
	fmt.Println(decodeString("40[leetcode]"))
}