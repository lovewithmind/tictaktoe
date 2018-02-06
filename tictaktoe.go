package main

import (
	"fmt"
)

func intializeArray(matrix *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = "X"
		}
	}
}

func printBoard(matrix [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func move(matrix *[3][3]string, x int, y int) {
	matrix[x][y] = "0"
}

func noWin(matrix [3][3]string) bool {
	return true
}

func main() {
	var matrix [3][3]string
	intializeArray(&matrix)

	fmt.Printf("Starting a new Game!!\n")
	fmt.Printf("Match Board:\n")

	printBoard(matrix)
	for noWin(matrix) {
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
