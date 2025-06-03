package console

import (
	"bufio"
	"strings"
	"testing"

	"github.com/KevinLiddle/tic_tac_go/core"
	"github.com/KevinLiddle/tic_tac_go/internal"
	"github.com/stretchr/testify/assert"
)

func BuildMockPlayerType(inputLines []string) (core.PlayerType, *strings.Builder) {
	var writer strings.Builder
	input := strings.Join(inputLines, "\n")
	inputScanner := bufio.NewScanner(strings.NewReader(input))

	playerType := core.PlayerType{
		Name:     "Mock",
		MakeMove: BuildConsoleHumanMover(inputScanner, &writer),
	}

	return playerType, &writer
}

func TestBuildConsoleHumanMover(t *testing.T) {
	t.Run("Prompts and accepts input for a move", func(t *testing.T) {
		input := []string{"3"}

		human, writer := BuildMockPlayerType(input)
		emptySpace := core.Player{}
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: human}
		playerO := core.Player{Token: "O", Color: internal.MOCK_GREEN, PlayerType: human}
		board := core.Board{
			[]core.Player{playerX, emptySpace, emptySpace},
			[]core.Player{playerO, playerX, emptySpace},
			[]core.Player{emptySpace, emptySpace, playerO},
		}

		assert.Equal(t, "", board[0][2].Token)
		human.MakeMove(board, playerX)

		expectedOutput := "\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n\n" +
			"  \033[1;REDmX\033[0m Choose an open space: "

		assert.Equal(t, expectedOutput, writer.String())
		assert.Equal(t, "X", board[0][2].Token)
	})

	t.Run("Re-prompts until a valid move is made", func(t *testing.T) {
		input := []string{"12", "3"}

		human, writer := BuildMockPlayerType(input)
		emptySpace := core.Player{}
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: human}
		playerO := core.Player{Token: "O", Color: internal.MOCK_GREEN, PlayerType: human}
		board := core.Board{
			[]core.Player{playerX, emptySpace, emptySpace},
			[]core.Player{playerO, playerX, emptySpace},
			[]core.Player{emptySpace, emptySpace, playerO},
		}

		assert.Equal(t, "", board[0][2].Token)
		human.MakeMove(board, playerO)

		expectedOutput := "\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n" +
			"\n" +
			"  \033[1;GREENmO\033[0m Choose an open space: " +
			"\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n" +
			"\n" +
			"  Invalid space. Try again.\n" +
			"  \n" +
			"  \033[1;GREENmO\033[0m Choose an open space: "

		assert.Equal(t, expectedOutput, writer.String())
		assert.Equal(t, "O", board[0][2].Token)
	})

	t.Run("Re-prompts until a free space is chosen", func(t *testing.T) {
		input := []string{"1", "4", "3"}

		human, writer := BuildMockPlayerType(input)
		emptySpace := core.Player{}
		playerX := core.Player{Token: "X", Color: internal.MOCK_RED, PlayerType: human}
		playerO := core.Player{Token: "O", Color: internal.MOCK_GREEN, PlayerType: human}
		board := core.Board{
			[]core.Player{playerX, emptySpace, emptySpace},
			[]core.Player{playerO, playerX, emptySpace},
			[]core.Player{emptySpace, emptySpace, playerO},
		}

		assert.Equal(t, "", board[0][2].Token)
		human.MakeMove(board, playerO)

		expectedOutput := "\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n" +
			"\n" +
			"  \033[1;GREENmO\033[0m Choose an open space: " +
			"\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n" +
			"\n" +
			"  That space is already taken. Try again.\n" +
			"  \n" +
			"  \033[1;GREENmO\033[0m Choose an open space: " +
			"\033c\n" +
			"     \033[1;REDmX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n" +
			"    -----------\n" +
			"     \033[1;GREENmO\033[0m | \033[1;REDmX\033[0m | \033[1;37m6\033[0m \n" +
			"    -----------\n" +
			"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;GREENmO\033[0m \n" +
			"\n" +
			"  That space is already taken. Try again.\n" +
			"  \n" +
			"  \033[1;GREENmO\033[0m Choose an open space: "

		assert.Equal(t, expectedOutput, writer.String())
		assert.Equal(t, "O", board[0][2].Token)
	})
}
