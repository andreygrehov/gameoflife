package life

type cellState int

const (
	dead cellState = iota
	alive
)

type cell int

func (c cell) String() string {
	if c == 0 {
		return " "
	}

	return "█"
}
