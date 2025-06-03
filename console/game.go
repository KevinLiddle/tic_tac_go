package console

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"

	"github.com/KevinLiddle/tic_tac_go/core"
)

type StringStringer string

func (s StringStringer) String() string {
	return string(s)
}

var playerColors = []TerminalColor{
	RED,
	GREEN,
	YELLOW,
	BLUE,
	MAGENTA,
}

var readyOptions = []StringStringer{"Let's play!", "I messed up."}
var endGameOptions = []StringStringer{"Rematch!", "New Game", "Quit"}

func getSelection[T fmt.Stringer](
	choices []T,
	prompt string,
	templateName string,
	templateData map[string]any,
	inputScanner *bufio.Scanner,
	writer io.Writer,
) T {
	var input string
	var inputInt int
	var err error

	templateData["Choices"] = choices
	templateData["Prompt"] = prompt

	for err != nil || inputInt <= 0 || inputInt > len(choices) {
		if len(input) > 0 && (err != nil || inputInt <= 0 || inputInt > len(choices)) {
			templateData["Error"] = "Invalid selection. Try again."
		}

		RenderTemplate(templateName, templateData, writer)

		inputScanner.Scan()
		input = inputScanner.Text()
		inputInt, err = strconv.Atoi(input)
	}

	return choices[inputInt-1]
}

func setupPlayers(inputScanner *bufio.Scanner, writer io.Writer) []core.Player {
	playerOptions := []core.PlayerType{
		core.PlayerType{
			Name:     "Human",
			MakeMove: BuildConsoleHumanMover(inputScanner, writer),
		},
	}
	players := make([]core.Player, 2)

	i := 0
	for i < len(players) {
		playerNumber := i + 1

		players[i].PlayerType = getSelection(
			playerOptions,
			fmt.Sprintf("Select a player type for Player %d:", playerNumber),
			"player_setup_selection",
			PlayersToTemplateData(players),
			inputScanner,
			writer,
		)

		var token string
		tokenTaken := func(p core.Player) bool { return p.Token == token }

		data := PlayersToTemplateData(players)
		data["Prompt"] = fmt.Sprintf("Select a token for Player %d (e.g. X): ", playerNumber)

		for len(token) != 1 || slices.ContainsFunc(players, tokenTaken) {
			if len(token) > 1 {
				data["Error"] = "Tokens must be a single character (e.g. X)"
			}

			if token != "" && slices.ContainsFunc(players, tokenTaken) {
				data["Error"] = fmt.Sprintf("%v is already taken. Choose another token.", token)
			}

			RenderTemplate("player_setup_text", data, writer)

			inputScanner.Scan()
			token = inputScanner.Text()
		}
		players[i].Token = token

		players[i].Color = getSelection(
			playerColors,
			fmt.Sprintf("Select a color for Player %d:", playerNumber),
			"player_setup_selection",
			PlayersToTemplateData(players),
			inputScanner,
			writer,
		)

		i++
	}
	return players
}

func setupGame(inputScanner *bufio.Scanner, writer io.Writer) []core.Player {
	var ready StringStringer
	var players []core.Player

	for ready != readyOptions[0] {
		players = setupPlayers(inputScanner, writer)
		ready = getSelection(
			readyOptions,
			"All good?",
			"player_setup_selection",
			PlayersToTemplateData(players),
			inputScanner,
			writer,
		)
	}

	return players
}

func takeTurns(players []core.Player) (core.Board, core.Player) {
	board := core.NewBoard(3)
	playerTurn := 0
	var winner core.Player

	for winner.IsEmpty() && !board.IsFull() {
		players[playerTurn].MakeMove(board)
		playerTurn = (playerTurn + 1) % len(players)
		winner = board.GetWinner()
	}

	return board, winner
}

func Play(reader io.Reader, writer io.Writer) {
	var players []core.Player
	inputScanner := bufio.NewScanner(reader)
	endGameChoice := endGameOptions[1]

	for endGameChoice != endGameOptions[2] {
		if endGameChoice == endGameOptions[1] {
			players = setupGame(inputScanner, writer)
		}

		board, winner := takeTurns(players)

		templateData := BoardToTemplateData(board)

		if !winner.IsEmpty() {
			templateData["Winner"] = PlayerToTemplateData(winner)
		}

		endGameChoice = getSelection(
			endGameOptions,
			"What next?",
			"end_game_selection",
			templateData,
			inputScanner,
			writer,
		)
	}
}
