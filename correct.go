package main

//IsCorrectRow Sudoku
func (s *Sudoku) IsCorrectRow(c [9][9]int, j int) bool {
	m := make(map[int]int)
	for i := 0; i < 9; i++ {
		num := c[i][j]
		if num == 0 {
			continue
		}
		m[num]++
		if m[num] > 1 {
			return false
		}
	}
	return true
}

//IsCorrectCol Sudoku
func (s *Sudoku) IsCorrectCol(c [9][9]int, i int) bool {
	m := make(map[int]int)
	for j := 0; j < 9; j++ {
		num := c[i][j]
		if num == 0 {
			continue
		}
		m[num]++
		if m[num] > 1 {
			return false
		}
	}
	return true
}

//IsCorrectBlock Sudoku
func (s *Sudoku) IsCorrectBlock(c [9][9]int, startI, startJ int) bool {
	m := make(map[int]int)
	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			num := c[i][j]
			if num == 0 {
				continue
			}
			m[num]++
			if m[num] > 1 {
				return false
			}
		}
	}
	return true
}

//IsCorrect Sudoku
func (s *Sudoku) IsCorrect(c [9][9]int) bool {
	for j := 0; j < 9; j++ {
		if !s.IsCorrectRow(c, j) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if !s.IsCorrectCol(c, i) {
			return false
		}
	}
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			if !s.IsCorrectBlock(c, i, j) {
				return false
			}
		}
	}
	return true
}
