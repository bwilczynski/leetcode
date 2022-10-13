package p27

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] != val {
			i++
			continue
		}
		for i < j-1 && nums[j] == val {
			j--
		}
		if nums[j] == val {
			break
		}
		nums[i] = nums[j]
		i++
		j--
	}
	return i
}
