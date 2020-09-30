package main



/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */
func isBadVersion(version int) bool{return true}
func firstBadVersion(n int) int {

	left := 1
	right := n
	for left+1 < right{
		mid := (right + left)/2
		if isBadVersion(mid) == true{
			right = mid
		}else if isBadVersion(mid) == false{
			left = mid
		}
	}
	if isBadVersion(left) == true{
		return left
	}
	if isBadVersion(right) == true{
		return right
	}
	return -1
}