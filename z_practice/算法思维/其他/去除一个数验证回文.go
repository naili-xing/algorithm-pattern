package 其他

func validPalindrome(s string) bool {
	slow := 0
	fast := len(s)-1

	for slow < fast{
		if s[slow] == s[fast] {
			slow +=1
			fast -=1
		}else{
			flat1, flat2 := true, true

			//跳过i
			for i,j :=slow, fast-1; i<j; i,j=i+1,j-1{
				if s[i] != s[j]{
					flat1 = false
					break
				}
			}
			//跳过j
			for i ,j :=slow+1, fast; i<j; i,j=i+1,j-1{
				if s[i] != s[j]{
					flat2 = false
					break
				}
			}
			return flat1 || flat2
		}
	}
	return true
}
