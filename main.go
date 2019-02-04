package main

import (
	"bufio"
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

var sudoku Sudoku

func init() {
	var gr [9][9]int
	var vr [9][9][9]bool
	sudoku = Sudoku{
		input:    gr,
		variants: vr,
		output:   gr,
	}
}

// PrintInput Sudoku
func (s *Sudoku) PrintInput() {
	fmt.Println("Input", s.input)
}

// PrintOutput Sudoku
func (s *Sudoku) PrintOutput() {
	fmt.Println("Output", s.output)
}

// LoadFromFile Sudoku
func (s *Sudoku) LoadFromFile(aFileName string) {
	s.ClearInput()

	_, err := ioutil.ReadFile(aFileName)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(aFileName)
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
		var str string = string(buffer)
		for j := 0; j < 9; j++ {
			var r2 int = int(str[j] - '0')
			if r2 == 240 {
				r2 = 0
			}
			s.input[i][j] = r2
		}
	}
	//s.Println()
}

// Resolve Sudoku
func (s *Sudoku) Resolve() {
	// init
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				if s.input[i][j] == k+1 {
					s.variants[i][j][k] = true
				} else {
					s.variants[i][j][k] = false
				}

			}
		}
	}
}

//ClearGrid Sudoku
func (s *Sudoku) ClearInput() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.input[i][j] = -1
		}
	}
}

func fromFile() {
	fmt.Println("Select sudoku file:")
	fmt.Println("(1) easy")
	fmt.Println("(2) normal")
	fmt.Println("(3) hard")
	fmt.Println("(4) veryhard")
	fmt.Println("(5) minimum")
	fmt.Println("(6) other")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	switch text[0] {
	case '1':
		{
			sudoku.LoadFromFile("schemes/easy.txt")
			sudoku.Resolve()
			sudoku.PrintInput()
			sudoku.PrintOutput()
		}
	case '2':
		{
			sudoku.LoadFromFile("schemes/normal.txt")
		}
	case '3':
		{
			sudoku.LoadFromFile("schemes/hard.txt")
		}
	case '4':
		{
			sudoku.LoadFromFile("schemes/veryhard.txt")
		}
	case '5':
		{
			sudoku.LoadFromFile("schemes/minimum.txt")
		}
	case '6':
		{
			fmt.Println("Write filename")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			sudoku.LoadFromFile("schemes/" + text)
		}
	default:
		fmt.Println("Wrong file " + text)
	}
}

func fromTerminal() {
	fmt.Println("fromTerminal")
}

func fromGenerate() {
	fmt.Println("fromGenerate")
}

func main() {
	fmt.Println("Select sudoku type:")
	fmt.Println("(1) from file")
	fmt.Println("(2) from terminal")
	fmt.Println("(3) from generate")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	switch text[0] {
	case '1':
		{
			fromFile()
		}
	case '2':
		{
			fromTerminal()
		}
	case '3':
		{
			fromGenerate()
		}
	default:
		fmt.Println("Wrong type " + text)
	}
}
