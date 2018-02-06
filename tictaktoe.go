package main

import (
	"fmt"
)

var win [8][3]int = [8][3]int{{0, 1, 2}, // Check first row.
	{3, 4, 5}, // Check second Row
	{6, 7, 8}, // Check third Row
	{0, 3, 6}, // Check first column
	{1, 4, 7}, // Check second Column
	{2, 5, 8}, // Check third Column
	{0, 4, 8}, // Check first Diagonal
	{2, 4, 6}} // Check second Diagonal

func intializeArray(matrix *[9]string) {
	for i := 0; i < 9; i++ {
		matrix[i] = "-"
	}
}

func printBoard(matrix [9]string) {
	for i := 0; i < 9; i += 3 {
		for j := i; j < i+3; j++ {
			fmt.Printf("%d{%s}\t", j+1, matrix[j])
		}
		fmt.Printf("\n")
	}
}

func move(matrix *[9]string, position int, sign string) {
	matrix[position-1] = sign
}

func isWin(matrix [9]string, c string) bool {
	for i := 0; i < 8; i++ {
		if matrix[win[i][0]] == c && matrix[win[i][1]] == c && matrix[win[i][2]] == c {
			return true
		}
	}
	return false
}

func main() {
	var matrix [9]string
	intializeArray(&matrix)
	playerOneSign := "X"
	playerTwoSign := "0"
	var x int
	moveCount := 0

	fmt.Printf("Starting a new Game!!\n")
	fmt.Printf("Match Board:\n")

	printBoard(matrix)
	for {

		fmt.Println("Player 1 move:")
		fmt.Scan(&x)
		move(&matrix, x, playerOneSign)
		printBoard(matrix)
		if isWin(matrix, playerOneSign) {
			fmt.Printf("Player 1 Won !!")
			break
		}
		moveCount++

		if moveCount == 9 {
			fmt.Printf("Game Draw !!\n")
			fmt.Printf("Starting a new Game:")
			intializeArray(&matrix)
			printBoard(matrix)
			continue
		}

		fmt.Println("Player 2 move:")
		fmt.Scan(&x)
		move(&matrix, x, playerTwoSign)
		printBoard(matrix)
		if isWin(matrix, playerTwoSign) {
			fmt.Printf("Player 2 Won !!")
			break
		}
		moveCount++
	}
}
