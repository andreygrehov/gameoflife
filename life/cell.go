package life

type cell int

const (
	_deadCell  = 0
	_aliveCell = 1
)

func (c cell) String() string {
	if c == _deadCell {
		return " "
	}

	return "█"
}
