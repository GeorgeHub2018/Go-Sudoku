package main

import "fmt"

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

//PrintInfo Sudoku
func (s *Sudoku) PrintInfo(a ...interface{}) {
	fmt.Println("INFO:", a)
}

//PrintWarning Sudoku
func (s *Sudoku) PrintWarning(a ...interface{}) {
	fmt.Println("WARNING:", a)
}

//PrintError Sudoku
func (s *Sudoku) PrintError(a ...interface{}) {
	fmt.Println("ERROR:", a)
}

//PrintIsInputCorrect Sudoku
func (s *Sudoku) PrintIsInputCorrect() bool {
	if !s.IsCorrect(s.input) {
		s.PrintError("input not correct")
		return false
	}
	s.PrintInfo("input correct")
	return true
}

//PrintIsOutputCorrect Sudoku
func (s *Sudoku) PrintIsOutputCorrect() bool {
	if !s.IsCorrect(s.output) {
		s.PrintError("output not correct")
		return false
	}
	s.PrintInfo("output correct")
	return true
}

//PrintHardLevel Sudoku
func (s *Sudoku) PrintHardLevel() {
	s.PrintInfo("hard level=", s.VariantsCount())
}
