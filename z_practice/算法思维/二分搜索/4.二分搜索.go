package main

func searchMatrix(matrix [][]int, target int) bool {

	left := 0
	right := len(matrix)*len(matrix[0])-1

	for left +1 < right{
		mid := (left +right)/2
		if matrix[mid/len(matrix[0])][mid%len(matrix[0])] < target{
			left = mid
		}else if matrix[mid/len(matrix[0])][mid%len(matrix[0])] > target {
			right = mid
		}else{
			return true
		}
	}

	if matrix[left/len(matrix[0])][left%len(matrix[0])] ==target{
		return true
	}
	if matrix[right/len(matrix[0])][right%len(matrix[0])] ==target{
		return true
	}

	return false
}

func main(){

}
