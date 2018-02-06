package main

import (
	"fmt"
)

func intializeArray(matrix *[9]string) {
	for i := 0; i < 9; i++ {
		matrix[i] = "X"
	}
}

func printBoard(matrix [9]string) {
	for i := 0; i < 9; i += 3 {
		for j := i; j < i+3; j++ {
			fmt.Printf("{" + matrix[j] + "}")
		}
		fmt.Printf("\n")
	}
}

func move(matrix *[9]string, x int, y int) {
	matrix[x] = "0"
}

func isWin(matrix [9]string, c string) bool {
	var win [8][3]int = [8][3]int{{0, 1, 2}, // Check first row.
		{3, 4, 5}, // Check second Row
		{6, 7, 8}, // Check third Row
		{0, 3, 6}, // Check first column
		{1, 4, 7}, // Check second Column
		{2, 5, 8}, // Check third Column
		{0, 4, 8}, // Check first Diagonal
		{2, 4, 6}} // Check second Diagonal

	for i := 0; i < 8; i++ {
		if matrix[win[i][0]] == c {
			return true
		}
	}
	return false
}

func main() {
	var matrix [9]string
	intializeArray(&matrix)

	fmt.Printf("Starting a new Game!!\n")
	fmt.Printf("Match Board:\n")

	printBoard(matrix)
	for !isWin(matrix, "") {
		fmt.Println("Player 1 move:")
		var x int
		var y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		move(&matrix, x, y)
		printBoard(matrix)
		fmt.Println("Player 2 move:")
		fmt.Scan(&x)
		fmt.Scan(&y)
		move(&matrix, x, y)
		printBoard(matrix)
	}
}
