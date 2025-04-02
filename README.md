# Peg Solitaire

A classic board game where the goal is to eliminate pegs until only one remains.

## How to Play

1. Choose a board size (7 or 9 recommended)
2. Move pegs (O) by jumping over adjacent pegs to remove them
3. To win, finish with only one peg left on the board

## Game Controls

- To select a peg: enter its position as `row,column`
- Choose direction: 0 (Up), 1 (Down), 2 (Left), 3 (Right)
- Type `B` to go back and correct the peg position

## Requirements

- Go (version 1.16 or later)

If you don't have Go installed, follow the official installation guide:
[https://golang.org/doc/install](https://golang.org/doc/install)

## Installation

```bash
# Clone the repository
git clone https://github.com/guischroeder/pego.git
cd pego

# Install dependencies
go mod tidy
```

## Running the Game

```bash
go run cmd/pego/main.go
```

## Example Board

```
  0 1 2 3 4 5 6
0     O O O    
1     O O O    
2 O O O O O O O
3 O O O · O O O
4 O O O O O O O
5     O O O    
6     O O O    
```

Legend:
- `O` = Peg
- `·` = Empty space

## Game Rules

1. Pegs can only jump over adjacent pegs into empty spaces
2. The jumped-over peg is removed from the board
3. Jumps are only allowed horizontally or vertically (not diagonally)
4. The game ends when no more jumps are possible
5. You win if only one peg remains

## Example Move

Starting with the board above, if you select the peg at position 5,3 and move it up (direction 0):
- The peg at 5,3 jumps over the peg at 4,3
- The peg at 4,3 is removed
- The peg lands at position 3,3 (which was empty)
