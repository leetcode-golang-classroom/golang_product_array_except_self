package sol

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	lower, upper := 1, 1
	for pos := 0; pos < n; pos++ {
		ans[pos] = 1
	}
	for pos := 0; pos < n; pos++ {
		if pos > 0 {
			lower *= nums[pos-1]
			upper *= nums[n-pos]
		}
		ans[pos] *= lower
		ans[n-1-pos] *= upper
	}
	return ans
}
