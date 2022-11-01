package p11

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		area := (j - i) * h
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
