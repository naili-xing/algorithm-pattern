package main

import "fmt"

func reverseBits(num uint32) uint32 {

	var ress uint32
	pointer := 31

	for num!=uint32(0){
		ress += (num&1) << pointer
		pointer -=1
		num = num >> 1
	}
	return ress
}



func main(){
	fmt.Println(reverseBits(00000010100101000001111010011100))
}
