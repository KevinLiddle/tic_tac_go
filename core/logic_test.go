package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockColor struct {
	name  string
	value string
}

func (c MockColor) Name() string {
	return c.name
}
func (c MockColor) Value() string {
	return c.value
}

var RED = MockColor{name: "Red", value: "red"}
var MockPlayer = PlayerType{
	Name:     "Mock",
	MakeMove: func(board Board, player Player) {},
}

func TestGetWinner(t *testing.T) {
	winners := []struct {
		board    [][]string
		expected string
	}{
		{
			[][]string{
				[]string{"", "", ""},
				[]string{"", "", ""},
				[]string{"", "", ""},
			},
			"",
		},
		{
			[][]string{
				[]string{"x", "o", "x"},
				[]string{"x", "o", "o"},
				[]string{"o", "x", "x"},
			},
			"",
		},
		{
			[][]string{
				[]string{"x", "x", "x"},
				[]string{"", "", ""},
				[]string{"", "", ""},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "", ""},
				[]string{"o", "o", "o"},
				[]string{"", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", ""},
				[]string{"", "", ""},
				[]string{"x", "x", "x"},
			},
			"x",
		},
		{
			[][]string{
				[]string{"x", "", ""},
				[]string{"x", "", ""},
				[]string{"x", "", ""},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "o", ""},
				[]string{"", "o", ""},
				[]string{"", "o", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "x"},
				[]string{"", "", "x"},
				[]string{"", "", "x"},
			},
			"x",
		},
		{
			[][]string{
				[]string{"x", "", ""},
				[]string{"", "x", ""},
				[]string{"", "", "x"},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "", "o"},
				[]string{"", "o", ""},
				[]string{"o", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"o", "o", "o", "o"},
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "", ""},
				[]string{"x", "x", "x", "x"},
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
				[]string{"o", "o", "o", "o"},
				[]string{"", "", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
				[]string{"", "", "", ""},
				[]string{"x", "x", "x", "x"},
			},
			"x",
		},
		{
			[][]string{
				[]string{"x", "", "", ""},
				[]string{"x", "", "", ""},
				[]string{"x", "", "", ""},
				[]string{"x", "", "", ""},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "o", "", ""},
				[]string{"", "o", "", ""},
				[]string{"", "o", "", ""},
				[]string{"", "o", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "o", ""},
				[]string{"", "", "o", ""},
				[]string{"", "", "o", ""},
				[]string{"", "", "o", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "", "x"},
				[]string{"", "", "", "x"},
				[]string{"", "", "", "x"},
				[]string{"", "", "", "x"},
			},
			"x",
		},
		{
			[][]string{
				[]string{"", "", "", "o"},
				[]string{"", "", "o", ""},
				[]string{"", "o", "", ""},
				[]string{"o", "", "", ""},
			},
			"o",
		},
		{
			[][]string{
				[]string{"", "", "", "j"},
				[]string{"", "", "j", ""},
				[]string{"", "j", "", ""},
				[]string{"j", "", "", ""},
			},
			"j",
		},
		{
			[][]string{
				[]string{"", "", "", ""},
				[]string{"", "", "j", ""},
				[]string{"", "j", "", ""},
				[]string{"j", "", "", ""},
			},
			"",
		},
	}

	for _, w := range winners {
		var rowBuilder strings.Builder
		for _, r := range w.board {
			for _, c := range r {
				if c == "" {
					rowBuilder.WriteString("-")
				} else {
					rowBuilder.WriteString(c)
				}
			}
		}

		t.Run(rowBuilder.String()+" winner is: "+w.expected, func(t *testing.T) {
			board := NewBoard(len(w.board))
			for row, rowValues := range w.board {
				for col, token := range rowValues {
					if token != "" {
						board[row][col] = Player{Token: token, Color: RED, PlayerType: MockPlayer}
					}
				}
			}
			assert.Equal(t, w.expected, GetWinner(board).Token)
		})
	}
}
