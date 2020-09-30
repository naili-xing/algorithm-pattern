package main

import "fmt"

func hammingWeight(num uint32) int {

	res := uint32(0)
	for i:=0; i<32; i++{
		res  += (num >> i) & uint32(1)

	}
	return int(res)
}


func  main(){

	fmt.Println(hammingWeight(00000000000000000000000000001011))

}
