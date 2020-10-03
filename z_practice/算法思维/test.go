package main
import "fmt"

func twoSum(nums []int, target int) []int {

	res := make(map[int]int)

	for i:=0;i<len(nums);i++{
		k := target-nums[i]
		res[k] = i
	}

	ress := []int{}

	for i:=0;i<len(nums);i++{
		k := nums[i]
		if _,ok:=res[k];ok{
			ress = append(ress, i)
			ress = append(ress, res[k])
			return ress
		}
	}
	return []int{}
}

func main() {
	fmt.Println(isValid("(])"))
}
