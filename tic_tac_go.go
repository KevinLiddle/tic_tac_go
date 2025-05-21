package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/KevinLiddle/tic_tac_go/core"
)

type StringStringer string

func (s StringStringer) String() string {
	return string(s)
}

type TerminalColor struct {
	name  string
	value string
}

func (c TerminalColor) Value() string {
	return c.value
}

func (c TerminalColor) Name() string {
	return c.name
}

func (c TerminalColor) String() string {
	return fmt.Sprintf("\033[%vm%s\033[0m", c.Value(), c.Name())
}

var RED = TerminalColor{name: "Red", value: "31"}
var GREEN = TerminalColor{name: "Green", value: "32"}
var YELLOW = TerminalColor{name: "Yellow", value: "33"}
var BLUE = TerminalColor{name: "Blue", value: "34"}
var MAGENTA = TerminalColor{name: "Magenta", value: "35"}
var WHITE = TerminalColor{name: "White", value: "37"}

var playerColors = []TerminalColor{
	RED,
	GREEN,
	YELLOW,
	BLUE,
	MAGENTA,
}

var readyOptions = []StringStringer{"Let's Play!", "I messed up."}
var endGameOptions = []StringStringer{"Rematch!", "New Game", "Quit"}

var playerOptions = []core.PlayerType{
	core.PlayerType{
		Name: "Human",
		MakeMove: func(board core.Board, player core.Player) {
			var moveInt int
			var rowIndex int
			var columnIndex int
			var err error

			inputScanner := bufio.NewScanner(os.Stdin)

			for err != nil || moveInt < 1 || moveInt > board.Size() || !board[rowIndex][columnIndex].IsEmpty() {
				printBoard(board)
				fmt.Printf("\033[%vm%v\033[0m Choose an open space: ", player.Color.Value(), player.Token)

				inputScanner.Scan()
				move := inputScanner.Text()
				moveInt, err = strconv.Atoi(move)
				rowIndex = (moveInt - 1) / board.Dimension()
				columnIndex = (moveInt - 1) % board.Dimension()
			}

			board[rowIndex][columnIndex] = player
		},
	},
}

func printBoard(board core.Board) {
	fmt.Printf("\033c")
	for rowIndex, row := range board {
		fmt.Printf("\t\n")
		for columnIndex, cell := range row {
			color := WHITE.Value()
			char := fmt.Sprintf("%d", columnIndex+(rowIndex*board.Dimension())+1)

			if !cell.IsEmpty() {
				color = cell.Color.Value()
				char = fmt.Sprintf("%v", cell.Token)
			}

			separator := ""
			if columnIndex < board.Dimension()-1 {
				separator = "|"
			}

			fmt.Printf("\033[1;%vm %v \033[0m%v", color, char, separator)
		}

		if rowIndex < board.Dimension()-1 {
			fmt.Printf("\t\n")

			for s := 0; s < 4*board.Dimension()-1; s++ {
				fmt.Printf("-")
			}
		}
	}
	fmt.Printf("\n\n")
}

func printTemplateLines(templateLines ...string) {
	for lineNumber, templateLine := range templateLines {
		newLine := "\n"
		if lineNumber == len(templateLines)-1 {
			newLine = ""
		}
		fmt.Printf("%v%v", templateLine, newLine)
	}
}

func printSetupTemplate(players []core.Player, templateLines ...string) {
	fmt.Printf("\033c")
	fmt.Printf("Welcome to Tic-Tac-Go(lang).\n\n")

	for i, player := range players {
		var color core.Color = WHITE
		if player.Color != nil {
			color = player.Color
		}

		fmt.Printf("Player %d: %v \033[%vm%s\033[0m\n", i+1, player.PlayerType, color.Value(), player.Token)
	}

	printTemplateLines(append([]string{"\n"}, templateLines...)...)
}

func getSelection[T fmt.Stringer](
	list []T,
	prompt string,
	printTemplate func(lines ...string),
	inputScanner *bufio.Scanner,
) T {
	var input string
	var inputInt int
	var err error

	for err != nil || inputInt <= 0 || inputInt > len(list) {
		output := []string{}

		for i, value := range list {
			output = append(output, fmt.Sprintf("\t%d) %v", i+1, value))
		}
		output = append(output, "")

		if input != "" && err != nil {
			output = append(output, "Invalid selection. Try again.")
		}
		output = append(output, prompt)

		printTemplate(output...)

		inputScanner.Scan()
		input = inputScanner.Text()
		inputInt, err = strconv.Atoi(input)
	}

	return list[inputInt-1]
}

func setupPlayers(inputScanner *bufio.Scanner) []core.Player {
	players := make([]core.Player, 2)

	i := 0
	for i < len(players) {
		playerNumber := i + 1

		players[i].PlayerType = getSelection(
			playerOptions,
			fmt.Sprintf("Select a player type for Player %d: ", playerNumber),
			func(lines ...string) { printSetupTemplate(players, lines...) },
			inputScanner,
		)

		var token string
		tokenTaken := func(p core.Player) bool { return p.Token == token }

		for len(token) != 1 || slices.ContainsFunc(players, tokenTaken) {
			output := []string{}

			if len(token) > 1 {
				output = append(output, "Tokens must be a single character (e.g. X)")
			}

			if token != "" && slices.ContainsFunc(players, tokenTaken) {
				output = append(output, fmt.Sprintf("%v is already taken. Choose another token.", token))
			}

			output = append(output, fmt.Sprintf("Select a token for Player %d (e.g. X): ", playerNumber))
			printSetupTemplate(players, output...)

			inputScanner.Scan()
			token = inputScanner.Text()
		}
		players[i].Token = token

		players[i].Color = getSelection(
			playerColors,
			fmt.Sprintf("Select a color for Player %d: ", playerNumber),
			func(lines ...string) { printSetupTemplate(players, lines...) },
			inputScanner,
		)

		i++
	}
	return players
}

func setupGame(inputScanner *bufio.Scanner) []core.Player {
	var ready StringStringer
	var players []core.Player

	for ready != readyOptions[0] {
		players = setupPlayers(inputScanner)
		ready = getSelection(
			readyOptions,
			"All good? ",
			func(lines ...string) { printSetupTemplate(players, lines...) },
			inputScanner,
		)
	}

	return players
}

func play(players []core.Player) (core.Board, core.Player) {
	board := core.NewBoard(3)
	playerTurn := 0
	var winner core.Player

	for winner.IsEmpty() {
		players[playerTurn].MakeMove(board)
		playerTurn = (playerTurn + 1) % len(players)
		winner = core.GetWinner(board)
	}

	return board, winner
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)

	var players []core.Player
	endGameChoice := endGameOptions[1]

	for endGameChoice != endGameOptions[2] {
		if endGameChoice == endGameOptions[1] {
			players = setupGame(inputScanner)
		}

		board, winner := play(players)
		endGameChoice = getSelection(
			endGameOptions,
			fmt.Sprintf("\033[%vm%v\033[0m Wins!\n\nWhat next? ", winner.Color.Value(), winner.Token),
			func(lines ...string) {
				printBoard(board)
				printTemplateLines(lines...)
			},
			inputScanner,
		)
	}
}
