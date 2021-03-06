package life

import (
	"strings"
)

const (
	boardRows = 20
	boardCols = 60
)

type board [boardRows][boardCols]cell

func (b board) String() string {
	var builder strings.Builder
	for i := range b {
		for j := range b[i] {
			builder.WriteString(b[i][j].String())
		}
		builder.WriteString(_lineBreak)
	}
	return builder.String()
}

func (b board) isAlive(i, j int) bool {
	return b[i][j] == _aliveCell
}

func (b board) isDead(i, j int) bool {
	return b[i][j] == _deadCell
}

func (b *board) setAlive(i, j int) {
	(*b)[i][j] = _aliveCell
}

func (b *board) setDead(i, j int) {
	(*b)[i][j] = _deadCell
}

func (b board) countNeighbours(i, j int) int {
	var neighbours int

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, 1},
		{1, 1}, {1, 0}, {1, -1}, {0, -1},
	}
	for _, dir := range dirs {
		ri := i + dir[0]
		rj := j + dir[1]

		if ri < 0 || ri == len(b) || rj < 0 || rj == len(b[i]) {
			continue
		}

		if b.isAlive(ri, rj) {
			neighbours++
		}
	}

	return neighbours
}
