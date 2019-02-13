# Go Sudoku Solver

![Image of gopher](https://raw.githubusercontent.com/GeorgeHub2018/Go-Sudoku/master/images/gopher.png)

**Go Sudoku** is small program written in **Golang** which can resolve sudoku.

There's three ways to load sudoku:
1. Load from **file** (example file or user file)
2. Load from **terminal** (by typing numbers)
3. Load from **example** which placed in program

## Sudoku file example

```
2 9     3
18   3 9 
64   725 
 9  61  8
  17 29  
5  94  6 
 753   19
 2 1   74
9     6 2
```

or

```
209000003
180003090
640007250
090061008
001702900
500940060
075300019
020100074
900000602
```

## How it works?

The algorithm of the program solves sudoku as well as a human would.
The program creates a matrix of possible values that are possible for each element and using the rules of non-repetition horizontally, vertically, and in a 3x3 square, by the elimination method, learn the necessary values. 
If the elimination method does not give the desired result, then the program finds a cell for which only 2 or 3 variants are possible and solves them by substituting possible values.
