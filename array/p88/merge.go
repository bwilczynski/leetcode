package p88

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 && m > 0 {
		l1, l2 := nums1[m-1], nums2[n-1]
		if l1 > l2 {
			nums1[m+n-1] = l1
			m--
		} else {
			nums1[m+n-1] = l2
			n--
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
