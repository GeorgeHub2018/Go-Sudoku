package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sudoku Sudoku

func init() {
	var ir [9][9]int
	var vr [9][9][9]bool
	var or [9][9]int
	sudoku = Sudoku{
		input:    ir,
		variants: vr,
		output:   or,
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
		}
	case '2':
		{
			sudoku.LoadFromFile("schemes/normal.txt")
			sudoku.Resolve()
		}
	case '3':
		{
			sudoku.LoadFromFile("schemes/hard.txt")
			sudoku.Resolve()
		}
	case '4':
		{
			sudoku.LoadFromFile("schemes/veryhard.txt")
			sudoku.Resolve()
		}
	case '5':
		{
			sudoku.LoadFromFile("schemes/minimum.txt")
			sudoku.Resolve()
		}
	case '6':
		{
			fmt.Println("Write file name")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			str := strings.Replace(text, `.\`, appDir()+`\`, -1)
			str = strings.Replace(str, `./`, appDir()+`/`, -1)
			str = strings.Replace(str, `\`, `/`, -1)
			if _, err := os.Stat(str); err == nil {
				sudoku.LoadFromFile(str)
				sudoku.Resolve()
			} else if os.IsNotExist(err) {
				fmt.Println("ERROR: file", str, "not exists")
				break
			} else {
				fmt.Println("ERROR:", err)
				break
			}
		}
	default:
		fmt.Println("Wrong file " + text)
	}
}

func fromTerminal() {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 9; i++ {
		fmt.Println("Input nine numbers [0, 9] (line ", i, ")")
		text, _ := reader.ReadString('\n')

		for j := 0; j < 9; j++ {
			r := int(text[j] - '0')
			if r == 240 {
				r = 0
			}
			if r == 221 {
				r = 0
			}
			sudoku.input[i][j] = r
		}
	}
	sudoku.SaveToFile(tempFileName(appDir()+"/temp/", ".txt"))
	sudoku.Resolve()
}

func fromExample() {
	sudoku.LoadFromInput(example)
	sudoku.Resolve()
}

func main() {
	fmt.Println("Select sudoku input:")
	fmt.Println("(1) from file")
	fmt.Println("(2) from terminal")
	fmt.Println("(3) from example")

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
			fromExample()
		}
	default:
		fmt.Println("Wrong type " + text)
	}
}
