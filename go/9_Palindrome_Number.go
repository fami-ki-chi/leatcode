func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	org := x
	rev := 0

	for x > 0 {
		tmp := x % 10
		x = x / 10
		rev = rev*10 + tmp
	}

	if rev == org {
		return true
	} else {
		return false
	}
}