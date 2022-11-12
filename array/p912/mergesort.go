package p912

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	m := len(nums) / 2
	l := sortArray(nums[:m])
	r := sortArray(nums[m:])
	return merge(l, r)
}

func merge(a1, a2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			res = append(res, a1[i])
			i++
		} else {
			res = append(res, a2[j])
			j++
		}
	}
	if i < len(a1) {
		res = append(res, a1[i:]...)
	} else if j < len(a2) {
		res = append(res, a2[j:]...)
	}
	return res
}
