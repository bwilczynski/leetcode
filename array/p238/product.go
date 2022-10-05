package p238

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	for i := 0; i < len(answer); i++ {
		answer[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1 - i
		answer[i] = left * answer[i]
		answer[j] = right * answer[j]
		left, right = nums[i]*left, nums[j]*right
	}
	return answer
}
