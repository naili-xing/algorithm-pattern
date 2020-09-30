package main

import "fmt"

/*
	3. Longest Substring Without Repeating Characters
	Given a string s, find the length of the longest substring without repeating characters.
	Example 1:
		Input: s = "abcabcbb"
		Output: 3
	Explanation: The answer is "abc", with the length of 3.
	Example 2:
		Input: s = "bbbbb"
		Output: 1
	Explanation: The answer is "b", with the length of 1.
	Example 3:
		Input: s = "pwwkew"
		Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
	Example 4:
	Input: s = ""
	Output: 0
 */

func lengthOfLongestSubstring(s string) int {
	if len(s)<=1{return len(s)}

	res := 0
	cn1 := [128]int{}
	c :=[]string{}
	i := 0
	j := 0
	for i<len(s) && j<len(s){
		//如果重复, i 右移，j不变
		if cn1[s[j]]>=1{
			cn1[s[i]]--
			i++
			c = c[1:]

			//如果不重复, j右移，i不变
		}else{
			cn1[s[j]]++
			c = append(c, string(s[j]))
			j++
		}
		if sum(cn1) >= res{
			res = sum(cn1)
		}
	}
	return res
}

func sum(l [128]int) int{
	res := 0
	for i:=0; i<len(l);i++{
		res +=l[i]
	}
	return res
}

func main(){
	fmt.Println(lengthOfLongestSubstring("  "))
}