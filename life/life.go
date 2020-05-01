package life

import (
	"fmt"
	"time"
)

const (
	boardRows = 20
	boardCols = 60
)

type Game struct {
	board [boardRows][boardCols]cell
}

type cellState int

const (
	dead cellState = iota
	alive
)

type cell int

func (c cell) String() string {
	if c.isDead() {
		return " "
	}

	return "█"
}

func (c cell) isAlive() bool {
	return cellState(c) == alive
}

func (c cell) isDead() bool {
	return cellState(c) == dead
}

func (c *cell) alive() {
	*c = 1
}

func (c *cell) dead() {
	*c = 0
}

func New() Game {
	return Game{
		board: [boardRows][boardCols]cell{
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
}

func (g *Game) step() {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	newBoard := g.board
	for i := range g.board {
		for j := range g.board[i] {
			var neighbours int
			oldCell := g.board[i][j]
			newCell := &newBoard[i][j]
			for _, dir := range dirs {
				ri := i + dir[0]
				rj := j + dir[1]

				if ri < 0 || ri == len(g.board) || rj < 0 || rj == len(g.board[i]) {
					continue
				}

				if g.board[ri][rj].isAlive() {
					neighbours++
				}
			}
			if oldCell.isAlive() && neighbours < 2 {
				newCell.dead()
				continue
			}
			if oldCell.isAlive() && (neighbours == 2 || neighbours == 3) {
				newCell.alive()
				continue
			}
			if oldCell.isAlive() && neighbours > 3 {
				newCell.dead()
				continue
			}
			if oldCell.isDead() && neighbours == 3 {
				newCell.alive()
				continue
			}
		}
	}
	g.board = newBoard
}

func (g *Game) render() {
	fmt.Print(_clearScreen)
	fmt.Print(_moveTo0x0)
	for i := range g.board {
		for j := range g.board[i] {
			c := g.board[i][j]
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (g *Game) Start() {
	for range time.Tick(200 * time.Millisecond) {
		g.render()
		g.step()
	}
}
