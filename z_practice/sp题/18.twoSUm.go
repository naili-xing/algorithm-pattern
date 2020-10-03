package main

import "fmt"

func twoSum(nums []int, target int) [][]int {

	res := make(map[int][]int)

	for i:=0;i<len(nums);i++{
		k := target-nums[i]
		if _,ok:=res[k];ok{

			res[k] = append(res[k], i)
		}else{
			res[k] = []int{}
			res[k] = append(res[k], i)
		}
	}
	var ress [][]int

	for i:=0;i<len(nums);i++{
		var tmp []int
		k := nums[i]
		if _,ok:=res[k];ok{
			for j:=0;j<len(res[k]);j++{
				if res[k][j]!=i{
					tmp = append(tmp, i)
					tmp = append(tmp,  res[k][j])
				}
			}
			ress = append(ress, tmp)
		}
	}
	return ress
}

func main(){
	fmt.Println(twoSum([]int{2,7,7,11,10}, 14))
}
