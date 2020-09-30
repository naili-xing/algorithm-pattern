package main

import "fmt"

/*
	给定一个 haystack 字符串和一个 needle 字符串，
	在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。

	输入: haystack = "hello", needle = "ll"
	输出: 2
	示例 2:

	输入: haystack = "aaaaa", needle = "bba"
	输出: -1

 */

func strStrTest(haystack string, needle string) int {


	if len(needle) <=0 {return 0}
	if len(needle)>len(haystack){return -1}


	for i:=0;i<=len(haystack)-len(needle);i++{

		for j:=0;j<len(needle);j++{
			if haystack[i+j] != needle[j]{
				break
			}
			if j==len(needle)-1{return i}
		}
	}
	return -1
}

func main(){
	fmt.Println(strStrTest("a", "a"))
}
