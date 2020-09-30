package main

import (
	"fmt"
	"sort"
)

/*
	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的数字可以无限制重复被选取。

	说明：

	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。 
	示例 1：

	输入：candidates = [2,3,6,7], target = 7,
	所求解集为：
	[
	  [7],
	  [2,2,3]
	]
	示例 2：

	输入：candidates = [2,3,5], target = 8,
	所求解集为：
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]
	 
	提示：
	1 <= candidates.length <= 30
	1 <= candidates[i] <= 200
	candidate 中的每个元素都是独一无二的。
	1 <= target <= 500

 */

func combinationSum(candidates []int, target int) [][]int {

	var res [][]int
	var list []int

	sort.Ints(candidates)
	combhelper(candidates, 0, target, list, &res)

	return res
}

func getSum(list []int)int{
	sv :=0
	for _, v := range list {
		sv += v
	}
	return sv
}

func combhelper(candidates []int, pos int, target int,  list []int, res *[][]int ){

	if 	sv := getSum(list); sv ==target{
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}
	for i:=0; i<len(candidates); i++{

		//递归上一层传入下一层pos,代表上一层检查的位置，下所有层操作完后，返回上一层，因为是值传递，所有回到上一层厚，pos还是原来的值
		if i<pos{continue}

		if 	sv := getSum(list);sv > target {continue}

		list = append(list, candidates[i])
		fmt.Println("befor ", pos, "checking=",candidates[i])
		combhelper(candidates, pos, target, list, res)
		fmt.Println("after ", pos, "checking=",candidates[i])
		pos +=1
		list = list[0: len(list)-1]
	}
}

func main(){

	fmt.Println(combinationSum([]int{2,3,6, 7}, 7))
}
