package main

import "fmt"

/*
	139. Word Break
		Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine
	if s can be segmented into a space-separated sequence of one or more dictionary words.
	Note:
		The same word in the dictionary may be reused multiple times in the segmentation.
		You may assume the dictionary does not contain duplicate words.
	Example 1:
		Input: s = "leetcode", wordDict = ["leet", "code"]
		Output: true
	Explanation: Return true because "leetcode" can be segmented as "leet code".
	Example 2:
		Input: s = "applepenapple", wordDict = ["apple", "pen"]
		Output: true
	Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
				 Note that you are allowed to reuse a dictionary word.
	Example 3:
		Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
		Output: false
 */


func wordBreak(s string, wordDict []string) bool {
	mapper := make(map[string]bool)

	for i:=0;i<len(wordDict);i++{
		mapper[wordDict[i]] = true
	}
	it := make([]bool, len(s)+1)
	it[0] = true
	for i:=1; i<len(it);i++{
		for j:=0;j<i;j++{
			if it[j] && mapper[s[j:i]] == true{
				it[i] = true
				break
			}
		}
	}
	return it[len(it)-1]
}

func main(){

	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))

}