package game

import (
	"pego/internal/board"
	"pego/internal/ui"
)

const (
	MinBoardSize = 5
	MaxBoardSize = 15
)

const (
	DirectionUp    = 0
	DirectionDown  = 1
	DirectionLeft  = 2
	DirectionRight = 3
)

type Game struct {
	board *board.Board
}

func NewGame() *Game {
	size := ui.GetBoardSize(MinBoardSize, MaxBoardSize)
	gameBoard := board.NewBoard(size)
	return &Game{board: gameBoard}
}

func (g *Game) Start() {
	for {
		ui.DisplayBoard(g.board.GetState(), board.EmptyCell, board.PegCell)

		if g.board.IsGameOver() {
			ui.DisplayGameOver(g.board.CountPegs())
			break
		}

		if !g.board.HasValidMoves() {
			ui.DisplayNoMovesAvailable()
			ui.DisplayGameOver(g.board.CountPegs())
			break
		}

		row, col, direction := ui.GetMove(g.board.Size())
		g.makeMove(row, col, direction)
	}
}

func (g *Game) makeMove(row, col, direction int) {
	var dx, dy int
	switch direction {
	case DirectionUp:
		dx, dy = -1, 0
	case DirectionDown:
		dx, dy = 1, 0
	case DirectionLeft:
		dx, dy = 0, -1
	case DirectionRight:
		dx, dy = 0, 1
	}

	g.board.MakeMove(row, col, dx, dy)
}
