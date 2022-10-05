package p1

func twoSum(nums []int, target int) []int {
	// a + b = target
	lookup := make(map[int]int)
	for i, a := range nums {
		b := target - a
		if ind, ok := lookup[b]; ok {
			return []int{ind, i}
		}
		lookup[a] = i
	}
	return []int{}
}
