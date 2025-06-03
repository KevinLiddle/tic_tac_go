package console

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/KevinLiddle/tic_tac_go/core"
)

func BuildConsoleHumanMover(
	inputScanner *bufio.Scanner,
	writer io.Writer,
) func(board core.Board, player core.Player) {
	return func(board core.Board, player core.Player) {
		var move string
		var moveInt int
		var rowIndex int
		var columnIndex int
		var err error

		for err != nil || moveInt < 1 || moveInt > board.Size() || !board[rowIndex][columnIndex].IsEmpty() {
			templateData := BoardToTemplateData(board)

			if len(move) > 0 {
				if err != nil || moveInt < 1 || moveInt > board.Size() {
					templateData["Error"] = "Invalid space. Try again."
				} else if !board[rowIndex][columnIndex].IsEmpty() {
					templateData["Error"] = "That space is already taken. Try again."
				}
			}
			templateData["Prompt"] = fmt.Sprintf(
				"%v Choose an open space: ",
				ColorizeBold(player.Color.Value(), player.Token),
			)

			RenderTemplate("human_move_prompt", templateData, writer)

			inputScanner.Scan()
			move = inputScanner.Text()
			moveInt, err = strconv.Atoi(move)
			rowIndex = (moveInt - 1) / board.Dimension()
			columnIndex = (moveInt - 1) % board.Dimension()
		}

		board[rowIndex][columnIndex] = player
	}
}
