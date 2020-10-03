package main

import (
	"fmt"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	res:=[][]int{}
	visited := make(map[int][][]int)

	for i :=0;i<R;i++{
		for j:=0;j<C;j++{
			if _, ok := visited[abs(i-r0)+ abs(j-c0)];ok{
				visited[abs(i-r0)+ abs(j-c0)] = append(visited[abs(i-r0)+ abs(j-c0)], []int{i,j})
			}else{
				visited[abs(i-r0)+ abs(j-c0)] = [][]int{}
				visited[abs(i-r0)+ abs(j-c0)] = append(visited[abs(i-r0)+ abs(j-c0)], []int{i,j})
			}
		}
	}

	j := 0
	for{
		if _,ok:=visited[j];ok{
			res = append(res, visited[j]...)
		}else{
			break
		}
		j+=1
	}

	return res
}

func abs(a int)int{
	if a>0{
		return a
	}else{
		return -a
	}
}

func main(){
	fmt.Println(allCellsDistOrder(2,2,0,1))
}
