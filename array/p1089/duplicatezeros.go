package p1089

func duplicateZeros(arr []int) {
	count0 := 0
	for _, v := range arr {
		if v == 0 {
			count0++
		}
	}
	if count0 == 0 {
		return
	}
	for i, j := len(arr)-1, len(arr)-1+count0; i >= 0; i-- {
		val := arr[i]
		if j < len(arr) {
			arr[j] = val
		}
		j--
		if val == 0 {
			if j >= 0 && j < len(arr) {
				arr[j] = 0
			}
			j--
		}
	}
}
