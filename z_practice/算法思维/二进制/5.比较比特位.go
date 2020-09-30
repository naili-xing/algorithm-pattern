package main

import "fmt"

func countBits(num int) []int {

	var ress []int

	for n:=0;n<=num;n++{
		res := 0
		for i:=0; i<32; i++{
			res  += (n >> i) & 1
		}
		ress = append(ress, res)
	}
	return ress
}


func main(){
	fmt.Println(countBits(5))
}
