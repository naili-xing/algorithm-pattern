package main

import "fmt"

/*
		1143. Longest Common Subsequence
		Given two strings text1 and text2, return the length of their longest common subsequence.

		A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.

		If there is no common subsequence, return 0.
		Example 1:

		Input: text1 = "abcde", text2 = "ace"
		Output: 3
		Explanation: The longest common subsequence is "ace" and its length is 3.
		Example 2:

		Input: text1 = "abc", text2 = "abc"
		Output: 3
		Explanation: The longest common subsequence is "abc" and its length is 3.
		Example 3:

		Input: text1 = "abc", text2 = "def"
		Output: 0
		Explanation: There is no such common subsequence, so the result is 0
 */

func longestCommonSubsequence(text1 string, text2 string) int {
	res := make([][]int, len(text1)+1)
	for i:=0;i<=len(text1);i++{
		res[i] = make([]int, len(text2)+1)
		for j:=0; j<=len(text2);j++{
			res[i][j] = 0
		}
	}
	for i:=1;i<len(res);i++{
		for j:=1; j<len(res[0]);j++{
			if text1[i-1] == text2[j-1]{
				res[i][j] = res[i-1][j-1]+1
			}else{
				res[i][j] = max2(res[i-1][j],res[i][j-1])
			}
		}
	}
	return res[len(text1)][len(text2)]
}

func max2(x,y int)int{
	if x>=y{
		return x
	}else{
		return y
	}
}

func main(){
 	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
