package console

import (
	"testing"

	"github.com/KevinLiddle/tic_tac_go/core"
	"github.com/KevinLiddle/tic_tac_go/internal"
	"github.com/stretchr/testify/assert"
)

func TestPlayerToTemplateData(t *testing.T) {
	t.Run("Turns a Player into a presentable object", func(t *testing.T) {
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: internal.MOCK_PLAYER_TYPE}
		expected := map[string]string{
			"Token":      "X",
			"ColorValue": "RED",
			"PlayerType": "MOCK",
		}
		assert.Equal(t, expected, PlayerToTemplateData(playerX))
	})
}

func TestPlayersToTemplateData(t *testing.T) {
	t.Run("Turns a list of Players into a presentable object", func(t *testing.T) {
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: internal.MOCK_PLAYER_TYPE}
		playerO := core.Player{Token: "O", Color: internal.MOCK_GREEN, PlayerType: internal.MOCK_PLAYER_TYPE}
		playerZ := core.Player{Token: "Z", Color: internal.MOCK_YELLOW, PlayerType: internal.MOCK_PLAYER_TYPE}
		expected := map[string]any{
			"Players": []map[string]any{
				{
					"Token":        "X",
					"ColorValue":   "RED",
					"PlayerType":   "MOCK",
					"PlayerNumber": 1,
				},
				{
					"Token":        "O",
					"ColorValue":   "GREEN",
					"PlayerType":   "MOCK",
					"PlayerNumber": 2,
				},
				{
					"Token":        "Z",
					"ColorValue":   "YELLOW",
					"PlayerType":   "MOCK",
					"PlayerNumber": 3,
				},
			},
		}
		assert.Equal(t, expected, PlayersToTemplateData([]core.Player{playerX, playerO, playerZ}))
	})
}

func TestBoardToTemplateData(t *testing.T) {
	t.Run("Turns a Board into a presentable object", func(t *testing.T) {
		emptySpace := core.Player{}
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: internal.MOCK_PLAYER_TYPE}
		playerO := core.Player{Token: "O", Color: internal.MOCK_GREEN, PlayerType: internal.MOCK_PLAYER_TYPE}
		board := core.Board{
			[]core.Player{playerX, emptySpace, emptySpace},
			[]core.Player{playerO, playerX, emptySpace},
			[]core.Player{emptySpace, emptySpace, playerO},
		}
		expected := map[string]any{
			"BoardDimension": 3,
			"Board": [][]map[string]string{
				{
					{
						"Token":      "X",
						"ColorValue": "RED",
						"PlayerType": "MOCK",
					},
					{
						"Token":      "2",
						"ColorValue": "37",
						"PlayerType": "",
					},
					{
						"Token":      "3",
						"ColorValue": "37",
						"PlayerType": "",
					},
				},
				{
					{
						"Token":      "O",
						"ColorValue": "GREEN",
						"PlayerType": "MOCK",
					},
					{
						"Token":      "X",
						"ColorValue": "RED",
						"PlayerType": "MOCK",
					},
					{
						"Token":      "6",
						"ColorValue": "37",
						"PlayerType": "",
					},
				},
				{
					{
						"Token":      "7",
						"ColorValue": "37",
						"PlayerType": "",
					},
					{
						"Token":      "8",
						"ColorValue": "37",
						"PlayerType": "",
					},
					{
						"Token":      "O",
						"ColorValue": "GREEN",
						"PlayerType": "MOCK",
					},
				},
			},
		}
		assert.Equal(t, expected, BoardToTemplateData(board))
	})
}
