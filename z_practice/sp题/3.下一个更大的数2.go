package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
556. 下一个更大元素 III
 */


func nextGreaterElement(n int) int {

	// 过程是，先从右边找到一个生序的组合
	// 在从这个组合右边找到一个比这个组合第一个数小的数， 交换
	// 右边升序排列
	if n>int(math.Pow(2,31)-1){return -1}
	sn := []byte(strconv.Itoa(n))
	for i:=len(sn)-2;i>=0;i--{
		if sn[i] < sn[i+1]{
				//找到了这个组合，找比他大的数
				for j:=len(sn)-1;j>i;j--{
					if sn[j]>sn[i]{
						tmp := sn[i]
						sn[i] = sn[j]
						sn[j] = tmp
						break
					}
				}
				//找到了这个组合，右边逆转
				for j:=i+1;j<=(len(sn)-i)/2+i;j++{
						tmp := sn[j]
						sn[j] = sn[len(sn)-j+i]
						sn[len(sn)-j+i] = tmp
				}
				v,_ := strconv.Atoi(string(sn[:]))
			if v>int(math.Pow(2,31)-1){return -1}

			return  v
			}
	}
	return -1
}

func main(){
	fmt.Println(nextGreaterElement(1999999999))//12223233
}