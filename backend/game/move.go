package game

//Move to lay a stone
type Move int

const (
	//MoveNone none
	MoveNone Move = iota
	//MoveCross turn for X
	MoveCross

	//MoveCircle turn for O
	MoveCircle
)

func (m Move) String() string {
	if m == MoveNone {
		return "_"
	}
	if m == MoveCross {
		return "X"
	}

	return "O"
}
