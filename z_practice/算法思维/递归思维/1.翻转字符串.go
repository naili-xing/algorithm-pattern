package main

/*
344. Reverse String
Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.



Example 1:

Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
 */

func reverseString(s []byte)  {
	i:=0
	j:=len(s)-1
	for i<j{
		s[i], s[j] = s[j], s[i]
		i+=1
		j-=1
	}

}


func main(){
	reverseString([]byte{'h','e','l','l','o'})
}
