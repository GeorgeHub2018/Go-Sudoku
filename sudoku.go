package main

import (
	"io/ioutil"
	"os"
)

type (
	// Sudoku struct
	Sudoku struct {
		input    [9][9]int
		variants [9][9][9]bool
		output   [9][9]int
	}
)

// LoadFromFile Sudoku
func (s *Sudoku) LoadFromFile(filename string) {
	s.ClearInput()

	_, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buffer := make([]byte, 11)
	for i := 0; i < 9; i++ {
		_, err := f.Read(buffer)
		if err != nil {
			panic(err)
		}
		str := string(buffer)
		for j := 0; j < 9; j++ {
			r2 := int(str[j] - '0')
			if r2 == 240 {
				r2 = 0
			}
			s.input[i][j] = r2
		}
	}
}

// LoadFromInput Sudoku
func (s *Sudoku) LoadFromInput(input [9][9]int) {
	s.input = input
}

//SaveToFile Sudoku
func (s *Sudoku) SaveToFile(filename string) {
	var str string
	for i := 0; i < 9; i++ {
		buf := make([]byte, 9)
		for j := 0; j < 9; j++ {
			buf[j] = byte(s.input[i][j]) + '0'
		}
		str = str + string(buf) + "\r\n"
		//f.WriteString(str)
	}
	writeStringToFile(filename, str)
	s.PrintInfo("file saved", filename)
}

//Clone Sudoku
func (s *Sudoku) Clone() Sudoku {
	var ir [9][9]int
	var vr [9][9][9]bool
	var or [9][9]int
	clone := Sudoku{
		input:    ir,
		variants: vr,
		output:   or,
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			clone.input[i][j] = s.input[i][j]
			clone.output[i][j] = s.output[i][j]
			for k := 0; k < 9; k++ {
				clone.variants[i][j][k] = s.variants[i][j][k]
			}
		}
	}
	return clone
}

//CellNum Sudoku
func (s *Sudoku) CellNum(i, j int) int {
	lastTrueIndex := 0
	trueCount := 0
	for k := 0; k < 9; k++ {
		if s.variants[i][j][k] == true {
			lastTrueIndex = k + 1
			trueCount++
		}
	}
	if trueCount > 1 {
		return 0
	}
	if trueCount == 9 {
		return 0
	}
	return lastTrueIndex
}

//VariantsCount Sudoku
func (s *Sudoku) VariantsCount() int {
	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cellNum := s.CellNum(i, j)
			if cellNum > 0 {
				count++
			}
		}
	}
	return count
}

//InitOutput Sudoku
func (s *Sudoku) InitOutput() {
	s.ClearOutput()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				cellNum := s.CellNum(i, j)
				if cellNum > 0 {
					s.output[i][j] = cellNum
				} else {
					s.output[i][j] = 0
				}
			}
		}
	}
}

//InitVariants Sudoku
func (s *Sudoku) InitVariants() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				if s.input[i][j] > 0 {
					if s.input[i][j] == k+1 {
						s.variants[i][j][k] = true
					} else {
						s.variants[i][j][k] = false
					}
				} else {
					s.variants[i][j][k] = true
				}
			}
		}
	}
}

// Resolve Sudoku
func (s *Sudoku) Resolve() {
	s.PrintInput()
	if !s.PrintIsInputCorrect() {
		return
	}

	s.InitVariants()
	s.PrintHardLevel()
	lastCount := 0
	repeatCount := 0
	for {
		s.ProcRows()
		s.ProcCols()
		s.ProcBlocks()
		count := s.VariantsCount()
		if count == 9*9 {
			s.PrintInfo("Resolved")
			break
		}
		if count == lastCount {
			repeatCount++
		} else {
			repeatCount = 0
		}
		if repeatCount >= 9 {
			s.PrintWarning("Not resolved")
			break
		}
		lastCount = count
	}
	s.InitOutput()
	s.PrintOutput()
	s.PrintIsOutputCorrect()
}

//ClearInput Sudoku
func (s *Sudoku) ClearInput() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.input[i][j] = -1
		}
	}
}

//ClearOutput Sudoku
func (s *Sudoku) ClearOutput() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.output[i][j] = 0
		}
	}
}
