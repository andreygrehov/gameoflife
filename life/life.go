package life

import (
	"fmt"
	"time"
)

const (
	_clearScreen = "\033[H\033[2J"
	_moveTo0x0   = "\033[0;0H"
)

type Game struct {
	activeBoard board
	shadowBoard board
}

func New() Game {
	g := Game{
		activeBoard: [boardRows][boardCols]cell{
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
	g.shadowBoard = g.activeBoard
	return g
}

func (g *Game) step() {
	for i := range g.activeBoard {
		for j := range g.activeBoard[i] {
			neighbours := g.activeBoard.countNeighbours(i, j)
			if g.activeBoard.isAlive(i, j) {
				if neighbours < 2 || neighbours > 3 {
					g.shadowBoard.setDead(i, j)
				}
			} else if neighbours == 3 {
				g.shadowBoard.setAlive(i, j)
			}
		}
	}
	g.activeBoard = g.shadowBoard
}

func (g *Game) render() {
	fmt.Print(_clearScreen)
	fmt.Print(_moveTo0x0)
	fmt.Println(g.activeBoard)
}

func (g *Game) Start() {
	for range time.Tick(200 * time.Millisecond) {
		g.render()
		g.step()
	}
}
