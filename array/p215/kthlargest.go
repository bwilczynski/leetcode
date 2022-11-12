package p215

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left, right, k int) int {
	p := partition(nums, left, right)
	if p == k {
		return nums[p]
	}
	if p > k {
		return quickSelect(nums, 0, p-1, k)
	} else {
		return quickSelect(nums, p+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	k := left
	for i := left; i < right; i++ {
		if nums[i] > nums[right] {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
	nums[k], nums[right] = nums[right], nums[k]
	return k
}
