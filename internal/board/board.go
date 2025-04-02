package board

const (
	EmptyCell = 0
	PegCell   = 1
)

const (
	DirectionUp    = 0
	DirectionDown  = 1
	DirectionLeft  = 2
	DirectionRight = 3
)

var directionVectors = map[int][2]int{
	DirectionUp:    {-1, 0},
	DirectionDown:  {1, 0},
	DirectionLeft:  {0, -1},
	DirectionRight: {0, 1},
}

type Board struct {
	cells [][]int
	size  int
}

func NewBoard(size int) *Board {
	board := &Board{
		cells: make([][]int, size),
		size:  size,
	}

	board.initialize()

	return board
}

func (b *Board) initialize() {
	for i := 0; i < b.size; i++ {
		b.cells[i] = make([]int, b.size)
		for j := 0; j < b.size; j++ {
			b.cells[i][j] = PegCell
		}
	}

	centerPos := b.size / 2
	b.cells[centerPos][centerPos] = EmptyCell
}

func (b *Board) GetState() [][]int {
	return b.cells
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) IsValidMove(row, col, direction int) bool {
	vector, exists := directionVectors[direction]
	if !exists {
		return false
	}

	dx, dy := vector[0], vector[1]

	if !b.isValidPosition(row, col) || b.cells[row][col] != PegCell {
		return false
	}

	jumpRow, jumpCol := row+dx, col+dy
	if !b.isValidPosition(jumpRow, jumpCol) || b.cells[jumpRow][jumpCol] != PegCell {
		return false
	}

	landRow, landCol := row+2*dx, col+2*dy
	if !b.isValidPosition(landRow, landCol) || b.cells[landRow][landCol] != EmptyCell {
		return false
	}

	return true
}

func (b *Board) isValidPosition(row, col int) bool {
	return row >= 0 && row < b.size && col >= 0 && col < b.size
}

func (b *Board) MakeMove(row, col, direction int) bool {
	if !b.IsValidMove(row, col, direction) {
		return false
	}

	vector := directionVectors[direction]
	dx, dy := vector[0], vector[1]

	jumpRow, jumpCol := row+dx, col+dy
	landRow, landCol := row+2*dx, col+2*dy

	b.movePeg(row, col, jumpRow, jumpCol, landRow, landCol)

	return true
}

func (b *Board) movePeg(startRow, startCol, jumpRow, jumpCol, landRow, landCol int) {
	b.cells[startRow][startCol] = EmptyCell
	b.cells[jumpRow][jumpCol] = EmptyCell
	b.cells[landRow][landCol] = PegCell
}

func (b *Board) HasValidMoves() bool {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.hasMoveFromPosition(i, j) {
				return true
			}
		}
	}

	return false
}

func (b *Board) hasMoveFromPosition(row, col int) bool {
	if b.cells[row][col] != PegCell {
		return false
	}

	for direction := DirectionUp; direction <= DirectionRight; direction++ {
		if b.IsValidMove(row, col, direction) {
			return true
		}
	}

	return false
}

func (b *Board) CountPegs() int {
	count := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.cells[i][j] == PegCell {
				count++
			}
		}
	}
	return count
}

func (b *Board) IsGameOver() bool {
	return b.CountPegs() == 1
}
