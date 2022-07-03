package gen

func binarySearch(seq []rune, target rune) int {
	lenSeq := len(seq)
	left := 0
	right := lenSeq - 1
	ans := right
	for left <= right {
		// left + right / 2
		mid := ((right - left) >> 1) + left
		if seq[mid] <= target {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
