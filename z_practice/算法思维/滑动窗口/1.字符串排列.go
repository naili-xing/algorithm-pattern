package main

import "fmt"

/*
	567. 字符串的排列
	给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

	换句话说，第一个字符串的排列之一是第二个字符串的子串。

	示例1:

	输入: s1 = "ab" s2 = "eidbaooo"
	输出: True
	解释: s2 包含 s1 的排列之一 ("ba").


	示例2:

	输入: s1= "ab" s2 = "eidboaoo"
	输出: False
 */


func checkInclusion(s1, s2 string) bool {

	if len(s1) > len(s2){
		return false
	}

	//26个字母，每个字母对应一个数字 index 为'b'-'a'
	cn1 := [26]int{}
	cn2 := [26]int{}

	for i:=0;i<len(s1);i++{
		cn1[s1[i]-'a']++
		cn2[s2[i]-'a']++
	}
	for i:=0;i<len(s2)-len(s1);i++{
		if cn1 == cn2{
			return true
		}
		// cn2减去第一个，在往后添加一个
		cn2[s2[i]-'a']--
		cn2[s2[i+len(s1)]-'a']++
	}
	return cn2 == cn1
}


func main(){
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}