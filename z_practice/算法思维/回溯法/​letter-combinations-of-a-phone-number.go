package main

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

输入："234"
输出：["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	var res []string
	if len(digits)==0 {return []string{}}
	letterHelper(digits, 0,"", &res)
	return res
}

func letterHelper(digits string, pos int, str string, res *[]string){
	if len(str) == len(digits){
		*res = append(*res, str)
	}
	if pos >= len(digits){return }
	//pos 代表了当前判断的数字是什么，只能从上往下判断，每个数字对应的每个跟节点，都要往下找
	for j:=0; j<len(phoneMap[string(digits[pos])]);j++{
		str += string(phoneMap[string(digits[pos])][j])
		letterHelper(digits, pos+1, str, res)
		str = str[0:len(str)-1]
	}
}
func main(){
	fmt.Println(letterCombinations("234"))
}


