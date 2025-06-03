package core

import (
	"slices"
)

type Color interface {
	Name() string
	Value() string
}

func ColorValue(c Color) string {
	if c == nil {
		return ""
	}
	return c.Value()
}

type PlayerType struct {
	Name     string
	MakeMove func(board Board, player Player)
}

func (pt PlayerType) String() string {
	return pt.Name
}

type Player struct {
	Token      string
	Color      Color
	PlayerType PlayerType
}

func (p Player) IsEmpty() bool {
	return p.Token == "" || p.Color.Value() == "" || p.PlayerType.Name == ""
}

func (p1 Player) Equals(p2 Player) bool {
	return p1.Token == p2.Token && ColorValue(p1.Color) == ColorValue(p2.Color)
}

func (player Player) MakeMove(board Board) {
	player.PlayerType.MakeMove(board, player)
}

type Board [][]Player

func (b Board) Dimension() int {
	return len(b)
}

func (b Board) Size() int {
	return len(b) * len(b)
}

func (b Board) IsFull() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell.IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (board Board) GetWinner() Player {
	var diag1 []Player
	var diag2 []Player

	for rowIndex := 0; rowIndex < board.Dimension(); rowIndex++ {
		if rowWinner := getWinnerFromLine(board[rowIndex]); !rowWinner.IsEmpty() {
			return rowWinner
		}

		var col []Player

		for colIndex := 0; colIndex < board.Dimension(); colIndex++ {
			col = append(col, board[colIndex][rowIndex])
		}

		if colWinner := getWinnerFromLine(col); !colWinner.IsEmpty() {
			return colWinner
		}

		diag1 = append(diag1, board[rowIndex][rowIndex])
		diag2 = append(diag2, board[board.Dimension()-rowIndex-1][rowIndex])
	}

	for _, diag := range [][]Player{diag1, diag2} {
		if diagWinner := getWinnerFromLine(diag); !diagWinner.IsEmpty() {
			return diagWinner
		}
	}
	return Player{}
}

func NewBoard(dimension int) Board {
	board := make([][]Player, dimension)
	for i := 0; i < dimension; i++ {
		board[i] = make([]Player, dimension)
	}
	return board
}

func getWinnerFromLine(line []Player) Player {
	compactedLine := slices.CompactFunc(
		slices.Clone(line),
		func(p1 Player, p2 Player) bool { return p1.Equals(p2) },
	)

	if len(compactedLine) == 1 {
		return compactedLine[0]
	}

	return Player{}
}
