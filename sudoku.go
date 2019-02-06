package main

import (
	"fmt"
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

// PrintInput Sudoku
func (s *Sudoku) PrintInput() {
	fmt.Println("Input ", s.input)
}

// PrintOutput Sudoku
func (s *Sudoku) PrintOutput() {
	fmt.Println("Output", s.output)
}

// PrintOutputVariants Sudoku
func (s *Sudoku) PrintOutputVariants() {
	fmt.Println("Variants", s.variants)
}

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
	fmt.Println("INFO: file saved", filename)
	//s.PrintInput()
}

//Clone Sudoku
func (s *Sudoku) Clone() Sudoku {
	var ir [9][9]int
	var vr [9][9][9]bool
	var or [9][9]int
	su := Sudoku{
		input:    ir,
		variants: vr,
		output:   or,
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			su.input[i][j] = s.input[i][j]
			su.output[i][j] = s.output[i][j]
			for k := 0; k < 9; k++ {
				su.variants[i][j][k] = s.variants[i][j][k]
			}
		}
	}
	return su
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

//MakeOutput Sudoku
func (s *Sudoku) MakeOutput() {
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

// Resolve Sudoku
func (s *Sudoku) Resolve() {
	s.InitVariants()
	fmt.Println("INFO: hard level=", s.VariantsCount())
	lastCount := 0
	repeatCount := 0
	for {
		s.ProcRows()
		s.ProcCols()
		s.ProcBlocks()
		count := s.VariantsCount()
		if count == 9*9 {
			fmt.Println("INFO: Resolved")
			break
		}
		if count == lastCount {
			repeatCount++
		} else {
			repeatCount = 0
		}
		if repeatCount >= 9 {
			fmt.Println("WARNING: Repeat limit")
			break
		}
		lastCount = count
	}
	s.MakeOutput()
	s.PrintInput()
	//s.PrintOutputVariants()
	s.PrintOutput()
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
