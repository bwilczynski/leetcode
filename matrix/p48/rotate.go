package p48

func rotate(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		for c := 0; c < len(matrix); c++ {
			matrix[i][c], matrix[j][c] = matrix[j][c], matrix[i][c]
		}
	}
	for i, row := range matrix {
		for j := i + 1; j < len(row); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
