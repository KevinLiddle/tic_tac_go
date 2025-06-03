package internal

import "github.com/KevinLiddle/tic_tac_go/core"

type MockColor struct {
	name  string
	value string
}

func NewMockColor(name string) MockColor {
	return MockColor{name: name, value: name}
}

func (c MockColor) Name() string {
	return c.name
}

func (c MockColor) Value() string {
	return c.value
}

var MOCK_RED = NewMockColor("RED")
var MOCK_GREEN = NewMockColor("GREEN")
var MOCK_YELLOW = NewMockColor("YELLOW")

var MOCK_PLAYER_TYPE = core.PlayerType{Name: "MOCK", MakeMove: func(b core.Board, p core.Player) {}}
