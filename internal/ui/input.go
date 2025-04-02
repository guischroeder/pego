package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"pego/internal/board"
)

func GetBoardSize(minSize, maxSize int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to Peg Solitaire!")
		fmt.Printf("Enter the board size (7 or 9 recommended):\n")

		input, _ := reader.ReadString('\n')
		rows, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || rows < minSize || rows > maxSize || rows%2 == 0 {
			fmt.Printf(ErrInvalidBoardSize, minSize, maxSize)
			continue
		}

		return rows
	}
}

func GetMove(boardSize int) (int, int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the row and column of the peg you want to move (e.g., 3,4):")
		input, _ := reader.ReadString('\n')

		row, col, valid := parseCoordinates(strings.TrimSpace(input), boardSize)
		if !valid {
			continue
		}

		direction := getDirection(reader)
		if direction == -1 {
			continue
		}

		return row, col, direction
	}
}

func parseCoordinates(input string, boardSize int) (int, int, bool) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		fmt.Println(ErrInvalidFormat)
		return 0, 0, false
	}

	row, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	col, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

	if err1 != nil || err2 != nil || !isValidPosition(row, col, boardSize) {
		fmt.Println(ErrInvalidPosition)
		return 0, 0, false
	}

	return row, col, true
}

func isValidPosition(row, col, boardSize int) bool {
	return row >= 0 && row < boardSize && col >= 0 && col < boardSize
}

func getDirection(reader *bufio.Reader) int {
	fmt.Println("Choose a direction:")
	fmt.Println("0: Up")
	fmt.Println("1: Down")
	fmt.Println("2: Left")
	fmt.Println("3: Right")

	input, _ := reader.ReadString('\n')
	direction, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || direction < board.DirectionUp || direction > board.DirectionRight {
		fmt.Println(ErrInvalidDirection)
		return -1
	}

	return direction
}
