func isValidSudoku(board [][]byte) bool {
    
	if len(board) != 9 {
		return false
	}
	for _, row := range board {
		if len(row) != 9 {
			return false
		}
	}

	for r := 0; r < 9; r++ {
		rowMap := make(map[byte]bool)

		for c := 0; c < 9; c++ {
			digit := board[r][c]
			if digit == '.' {
				continue
			}
			if rowMap[digit] {
				return false
			}
			rowMap[digit] = true
		}
	}

	for c := 0; c < 9; c++ {
		columnMap := make(map[byte]bool)

		for r := 0; r < 9; r++ {
			digit := board[r][c]
			if digit == '.' {
				continue
			}
			if columnMap[digit] {
				return false
			}
			columnMap[digit] = true
		}
	}


	var  hasDuplicates func(r,c int) bool

	hasDuplicates = func(r,c int) bool{
        seen := make(map[byte]bool)

		for i:= r ; i < r+3 ; i++{
            for j:=c; j< c+3; j++{
                if board[i][j] == '.' {
					continue
				}
				if seen[board[i][j]]{
					return true
				}

				seen[board[i][j]] = true
			}
		}
		return false
	}


	for r:=0 ; r<9 ; r=r+3{
		for c:=0; c<9 ; c=c+3{
			if hasDuplicates(r,c){
                 return false
			}
		}
	}
    return true
}