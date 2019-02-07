package main

//ProcRow Sudoku
func (s *Sudoku) ProcRow(j int) {
	// proc by num
	for i := 0; i < 9; i++ {
		cellNum := s.CellNum(i, j)
		if cellNum > 0 {
			for ii := 0; ii < 9; ii++ {
				if i == ii {
					continue
				}
				s.variants[ii][j][cellNum-1] = false
			}
		}
	}
	// proc by empty
	for num := 1; num <= 9; num++ {
		numCount := 0
		lastNumIndex := 0
		for i := 0; i < 9; i++ {
			if s.variants[i][j][num-1] == true {
				numCount++
				lastNumIndex = i
			}
		}
		if numCount == 1 {
			for k := 0; k < 9; k++ {
				if k == num-1 {
					continue
				}
				s.variants[lastNumIndex][j][k] = false
			}
		}
	}
}

//ProcRows Sudoku
func (s *Sudoku) ProcRows() {
	for j := 0; j < 9; j++ {
		s.ProcRow(j)
	}
}

//ProcCol Sudoku
func (s *Sudoku) ProcCol(i int) {
	// proc by num
	for j := 0; j < 9; j++ {
		cellNum := s.CellNum(i, j)
		if cellNum > 0 {
			for jj := 0; jj < 9; jj++ {
				if j == jj {
					continue
				}
				s.variants[i][jj][cellNum-1] = false
			}
		}
	}
	// proc by empty
	for num := 1; num <= 9; num++ {
		numCount := 0
		lastNumIndex := 0
		for j := 0; j < 9; j++ {
			if s.variants[i][j][num-1] == true {
				numCount++
				lastNumIndex = j
			}
		}
		if numCount == 1 {
			for k := 0; k < 9; k++ {
				if k == num-1 {
					continue
				}
				s.variants[i][lastNumIndex][k] = false
			}
		}
	}
}

//ProcCols Sudoku
func (s *Sudoku) ProcCols() {
	for i := 0; i < 9; i++ {
		s.ProcCol(i)
	}
}

//ProcBlock Sudoku
func (s *Sudoku) ProcBlock(startI, startJ int) {
	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			cellNum := s.CellNum(i, j)
			if cellNum > 0 {
				for ii := startI; ii < startI+3; ii++ {
					for jj := startJ; jj < startJ+3; jj++ {
						if (i == ii) && (j == jj) {
							continue
						}
						s.variants[ii][jj][cellNum-1] = false
					}
				}
			}
		}
	}

	// proc by empty
	for num := 1; num <= 9; num++ {
		numCount := 0
		lastCol := 0
		lastRow := 0
		for i := startI; i < startI+3; i++ {
			for j := startJ; j < startJ+3; j++ {
				if s.variants[i][j][num-1] == true {
					numCount++
					lastCol = i
					lastRow = j
				}
			}
		}
		if numCount == 1 {
			for k := 0; k < 9; k++ {
				if k == num-1 {
					continue
				}
				s.variants[lastCol][lastRow][k] = false
			}
		}
	}
}

//ProcBlocks Sudoku
func (s *Sudoku) ProcBlocks() {
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			s.ProcBlock(i, j)
		}
	}
}
