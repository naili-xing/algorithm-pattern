package main

import "fmt"

/*
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
		Example 1:
			Input: "babad"
			Output: "bab"
			Note: "aba" is also a valid answer.
		Example 2:
			Input: "cbbd"
			Output: "bb"
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/longest-palindromic-substring
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func longestPalindrome(s string) string {

	res := s
	minv :=0
	for i:=0;i<len(s)-1;i++{
		for j:=i+1;j<=len(s);j++{
			if isPal(s[i:j]) && len(s[i:j])>minv{
				res =  s[i:j]
				minv = len(s[i:j])
			}
		}
	}
	return res
}

func isPal(s string)bool{
	i:=0
	j:=len(s)-1
	for i<j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}


func main(){

	fmt.Println(longestPalindrome("bb"))

}



