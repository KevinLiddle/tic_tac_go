package console

import (
	"fmt"

	"github.com/KevinLiddle/tic_tac_go/core"
)

func PlayerToTemplateData(player core.Player) map[string]string {
	var color core.Color = WHITE
	if player.Color != nil {
		color = player.Color
	}

	return map[string]string{
		"PlayerType": player.PlayerType.String(),
		"ColorValue": color.Value(),
		"Token":      player.Token,
	}
}

func PlayersToTemplateData(players []core.Player) map[string]any {
	playerTemplates := make([]map[string]any, len(players))

	for i, player := range players {
		p := make(map[string]any)
		for k, v := range PlayerToTemplateData(player) {
			p[k] = v
		}
		p["PlayerNumber"] = i + 1
		playerTemplates[i] = p
	}
	return map[string]any{"Players": playerTemplates}
}

func BoardToTemplateData(board core.Board) map[string]any {
	templateData := make(map[string]any)
	templateBoard := make([][]map[string]string, board.Dimension())

	for rowIndex, row := range board {
		templateBoard[rowIndex] = make([]map[string]string, board.Dimension())
		for colIndex, col := range row {
			templateBoard[rowIndex][colIndex] = PlayerToTemplateData(col)
			if len(templateBoard[rowIndex][colIndex]["Token"]) == 0 {
				templateBoard[rowIndex][colIndex]["Token"] = fmt.Sprintf("%d", rowIndex*board.Dimension()+colIndex+1)
			}
		}
	}
	templateData["Board"] = templateBoard
	templateData["BoardDimension"] = board.Dimension()

	return templateData
}
