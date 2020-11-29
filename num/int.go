package num

// Sign of a integer
// 1 for positive
// 0 for zero
// -1 for negative
func Sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return +1
	}
	return 0
}

// MaxInt is compares and returns the larger one.
func MaxInt(first int, rest ...int) int {
	ans := first
	for _, item := range rest {
		if item > ans {
			ans = item
		}
	}
	return ans
}

// MinInt is compares and returns the lesser one.
func MinInt(first int, rest ...int) int {
	ans := first
	for _, item := range rest {
		if item < ans {
			ans = item
		}
	}
	return ans
}
