package ui

import (
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayBoard(gameBoard [][]int, emptyCell, pegCell int) {
	clearScreen()

	rows := len(gameBoard)
	cols := len(gameBoard[0])

	printColumnHeaders(cols)
	printBoardRows(gameBoard, rows, cols)
	fmt.Println()
}

func printColumnHeaders(cols int) {
	fmt.Print("  ")
	for j := 0; j < cols; j++ {
		fmt.Printf(" %d", j)
	}
	fmt.Println()
}

func printBoardRows(gameBoard [][]int, rows, cols int) {
	for i := 0; i < rows; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < cols; j++ {
			printCell(gameBoard[i][j])
		}
		fmt.Println()
	}
}

func printCell(cellValue int) {
	switch cellValue {
	case 0: // EmptyCell
		fmt.Printf(" %s", EmptyCellSymbol)
	case 1: // PegCell
		fmt.Printf(" %s", PegCellSymbol)
	default:
		fmt.Print("  ")
	}
}

func DisplayGameOver(pegsRemaining int) {
	fmt.Println("\nGame Over!")
	fmt.Printf("Pegs remaining: %d\n", pegsRemaining)

	if pegsRemaining == 1 {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Try again to get down to just one peg!")
	}
}

func DisplayNoMovesAvailable() {
	fmt.Println("\nNo more moves available!")
}
