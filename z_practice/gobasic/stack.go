package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strStr(haystack string, needle string) int{

	if len(haystack) == 0{return 0}

	for i:=0; i<=len(haystack) - len(needle); i++{
		for j := range needle{
			if haystack[j+i] != needle[j]{
				break
			}
			if len(needle)-1 == j{
				return i
			}
		}
	}
	return -1
}

func ihash(s string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}


func main() {

	// stack
	stack := make([]int, 0)
	stack = append(stack, 1,2,3)
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("Stack", stack,len(stack), stack[0], v)

	 // 队列
	 queue := make([]int, 0)
	 queue = append(queue, 10,100,20,30)
	 v2 := queue[0]
	 queue = queue[1:]
	 fmt.Println("Queue",queue, v2,len(queue) == 0)

	 //dict
	 m := make(map[string]int)
	 m["a1"] =1
	 m["a2"] =1
	 delete(m, "a1")
	 for k, v := range m {
	 	println(k, " ",v)
	 }
	 fmt.Println(m)

	 // sort
	 sort.Ints([]int{1,2,3,9,1})
	 sort.Strings([]string{"a", "b"})
	 s := []int{1,2,3,4,5,6,7}
	 sort.Slice(s, func(i,j int)bool{return s[i]<s[j]})

	 //math
	 fmt.Println(math.MaxInt32)

	 // copy
	 a := []int{1,2,3,4,5}
	 // delete a[1]
	 copy(a[1:], a[1+1:])
	 a = a[:len(a)-1]
	 fmt.Println(a)

	 b := make([]int, 10)
	 b[1] = 123
	 fmt.Println(b)

	 s2 := "1234"
	 num := int(s2[0] - '0')
	 fmt.Println(num)

	 b3 := byte(num+'0')
	 fmt.Println(b3)
	 words := ihash("asdfd")
	 fmt.Println(words)
	type KeyValue struct {
		Key   string
		Value string
	}
	fileName := "test.txt"
	var pairs KeyValue

	filteredData := map[string][]string{}
	fmt.Println("2. read file", fileName)
	newFile, _ := os.Open(fileName)
	decoder := json.NewDecoder(newFile)
	for decoder.More() {
		err := decoder.Decode(&pairs)
		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Println("3. data is", pairs)

			if _, ok := filteredData[pairs.Key]; !ok {
				filteredData[pairs.Key] = []string{}
			}
			filteredData[pairs.Key] = append(filteredData[pairs.Key], pairs.Key)
		}
			fmt.Println(filteredData)
	}

	var ress int
	for _, s := range []string{"123", "444"}{
		//res += strconv.Atoi(e)
		sint, _ := strconv.Atoi(s)
		ress += sint
	}
	fmt.Println(strconv.Itoa(123))
	//return string(res)

	for {
		fmt.Println(strings.ToLower("Gopher"))
		break
	}

	fmt.Println("num of task in each worker is ",7/2)
}
