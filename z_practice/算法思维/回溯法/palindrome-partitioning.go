package main

import "fmt"

/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。
	示例:
		输入: "aab"
		输出:
		[
		  ["aa","b"],
		  ["a","a","b"]
		]
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/palindrome-partitioning
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func partition(s string) [][]string {

	var res [][]string
	var list []string
	parHelper(s, list, &res, s)
	return res
}
func isPar(s string)bool{
	for i :=0; i<len(s)/2; i++{
		if s[i]!=s[len(s)-i-1]{
			return false
		}
	}
	return true
}

func parHelper(s string, list []string, res *[][]string, gs string){

	ts := ""
	for _,v := range list{
		ts+=v
	}
	if len(ts) == len(gs) {
			ans := make([]string, len(list))
			copy(ans, list)
			*res = append(*res, ans)
		}
	for i:=0; i<len(s); i++{
		if !isPar(s[0:i+1]){continue}
		if i+1>len(s){return }
		list = append(list, s[0:i+1])
		parHelper(s[i+1:], list, res, gs)
		list = list[0:len(list)-1]
	}
}

func main(){
	fmt.Println(partition("cbbbcc"))

}
