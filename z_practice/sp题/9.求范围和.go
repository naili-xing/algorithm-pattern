package main


type NumArray struct {

	res []int

}

func Constructor(nums []int) NumArray {
	return NumArray{nums}

}


func (this *NumArray) SumRange(i int, j int) int {
	r:=0
	for k:=i;k<=j;k++{
		r+=this.res[k]
	}
	return r
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main(){

}
