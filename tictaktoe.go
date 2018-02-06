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

func main() {
	var matrix [3][3]string
	intializeArray(&matrix)

	fmt.Printf("Starting a new Game!!\n")
	fmt.Printf("Match Board:\n")

	printBoard(matrix)

	fmt.Println("Player 1 move")
}
