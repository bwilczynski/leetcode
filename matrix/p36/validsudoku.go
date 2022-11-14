package p36

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, len(board))
	cols := make([]map[byte]struct{}, len(board))
	boxes := make([]map[byte]struct{}, len(board))

	boxIndex := func(i, j int) int {
		return 3*(i/3) + j/3
	}

	for i := 0; i < len(board); i++ {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		boxes[i] = make(map[byte]struct{})
	}
	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			if _, ok := rows[i][val]; ok {
				return false
			}
			if _, ok := cols[j][val]; ok {
				return false
			}
			k := boxIndex(i, j)
			if _, ok := boxes[k][val]; ok {
				return false
			}
			rows[i][val] = struct{}{}
			cols[j][val] = struct{}{}
			boxes[k][val] = struct{}{}
		}
	}
	return true
}
