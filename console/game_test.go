package console

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	t.Run("Plays a Human vs Human game - X Winner", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"X", // Player 1 token
			"1", // Player 1 color: Red
			"1", // Player 2 type: Human
			"O", // Player 2 token
			"2", // Player 2 color: Green
			"1", // Confirm
			"1", // X
			"2", // O
			"4", // X
			"5", // O
			"7", // X
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[0],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[1],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mX\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[3],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[4],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[37mO\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[32mO\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[6],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mX\033[0m Choose an open space: ",
			actualOutput[7],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mX\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mO\033[0m Choose an open space: ",
			actualOutput[8],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mX\033[0m Choose an open space: ",
			actualOutput[9],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mX\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mO\033[0m Choose an open space: ",
			actualOutput[10],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mX\033[0m Choose an open space: ",
			actualOutput[11],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mX\033[0m | \033[1;32mO\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mX\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[31mX\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[12],
		)
	})

	t.Run("Plays a Human vs Human game - Cat's Game", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"Y", // Player 1 token
			"3", // Player 1 color: Yellow
			"1", // Player 2 type: Human
			"B", // Player 2 token
			"4", // Player 2 color: Blue
			"1", // Confirm
			"1", // Y
			"2", // B
			"4", // Y
			"5", // B
			"8", // Y
			"7", // B
			"3", // Y
			"6", // B
			"9", // Y
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[0],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[1],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[3],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[4],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37mB\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[34mB\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[6],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[7],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[8],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[9],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[10],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[11],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;33mY\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[12],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;33mY\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[13],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;33mY\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;33mY\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[14],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;33mY\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;34mB\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;33mY\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[15],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;33mY\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;34mB\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;33mY\033[0m | \033[1;33mY\033[0m \n\n"+
				"Cat's Game...       (ノಠ益ಠ)ノ彡┻━┻\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[16],
		)
	})

	t.Run("Plays a Human vs Human game - Invalid token length", func(t *testing.T) {
		inputLines := []string{
			"1",  // Player 1 type: Human
			"YY", // Player 1 token error
			"Y",  // Player 1 token
			"1",  // Player 1 color: Red
			"1",  // Player 2 type: Human
			"U",  // Player 2 token
			"2",  // Player 2 color: Green
			"1",  // Confirm
			"1",  // Y
			"2",  // U
			"4",  // Y
			"5",  // U
			"7",  // Y
			"3",  // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Tokens must be a single character (e.g. X)\n"+
				"  \n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[3],
		)
	})

	t.Run("Plays a Human vs Human game - Token already used", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"Y", // Player 1 token
			"1", // Player 1 color: Red
			"1", // Player 2 type: Human
			"Y", // Player 2 token
			"U", // Player 2 token
			"2", // Player 2 color: Green
			"1", // Confirm
			"1", // Y
			"2", // U
			"4", // Y
			"5", // U
			"7", // Y
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Y is already taken. Choose another token.\n"+
				"  \n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2: Human \033[37mU\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[6],
		)
	})

	t.Run("Plays a Human vs Human game - Messed up the setup", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"X", // Player 1 token
			"1", // Player 1 color: Red
			"1", // Player 2 type: Human
			"O", // Player 2 token
			"2", // Player 2 color: Green
			"2", // I messed up.
			"1", // Player 1 type: Human
			"Y", // Player 1 token
			"1", // Player 1 color: Red
			"1", // Player 2 type: Human
			"U", // Player 2 token
			"2", // Player 2 color: Green
			"1", // Confirm
			"1", // Y
			"2", // U
			"4", // Y
			"5", // U
			"7", // Y
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[0],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[1],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mX\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[3],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[4],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[37mO\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mX\033[0m\n"+
				"  Player 2: Human \033[32mO\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[6],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[7],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[8],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[9],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[10],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[11],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2: Human \033[37mU\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[12],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[31mY\033[0m\n"+
				"  Player 2: Human \033[32mU\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[13],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mY\033[0m Choose an open space: ",
			actualOutput[14],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mU\033[0m Choose an open space: ",
			actualOutput[15],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mY\033[0m Choose an open space: ",
			actualOutput[16],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mY\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mU\033[0m Choose an open space: ",
			actualOutput[17],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;31mY\033[0m Choose an open space: ",
			actualOutput[18],
		)
		assert.Equal(
			t,
			"\n     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mY\033[0m | \033[1;32mU\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;31mY\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[31mY\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[19],
		)
	})

	t.Run("Plays a Human vs Human game - Rematch", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"Y", // Player 1 token
			"3", // Player 1 color: Yellow
			"1", // Player 2 type: Human
			"B", // Player 2 token
			"4", // Player 2 color: Blue
			"1", // Confirm
			"1", // Y
			"2", // B
			"4", // Y
			"5", // B
			"7", // Y
			"1", // Rematch
			"1", // Y
			"5", // B
			"4", // Y
			"7", // B
			"2", // Y
			"3", // B
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[0],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[1],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[3],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[4],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37mB\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[34mB\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[6],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[7],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[8],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[9],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[10],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[11],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[33mY\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[12],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[13],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[14],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[15],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[16],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[17],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;33mY\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[18],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;33mY\033[0m | \033[1;34mB\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;34mB\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[34mB\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[19],
		)
	})

	t.Run("Plays a Human vs Human game - New Game", func(t *testing.T) {
		inputLines := []string{
			"1", // Player 1 type: Human
			"Y", // Player 1 token
			"3", // Player 1 color: Yellow
			"1", // Player 2 type: Human
			"B", // Player 2 token
			"4", // Player 2 color: Blue
			"1", // Confirm
			"1", // Y
			"2", // B
			"4", // Y
			"5", // B
			"7", // Y
			"2", // New Game
			"1", // Player 1 type: Human
			"G", // Player 1 token
			"2", // Player 1 color: Green
			"1", // Player 2 type: Human
			"M", // Player 2 token
			"5", // Player 2 color: Magenta
			"1", // Confirm
			"1", // G
			"5", // M
			"4", // G
			"7", // M
			"2", // G
			"3", // M
			"3", // Quit
		}
		var writer strings.Builder
		reader := strings.NewReader(strings.Join(inputLines, "\n"))

		Play(reader, &writer)

		// Skip the first item because we clear the screen initially
		actualOutput := strings.Split(writer.String(), "\033c")[1:]
		assert.Equal(t, len(inputLines), len(actualOutput))
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[0],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[1],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[2],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[3],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[4],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[37mB\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[5],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[33mY\033[0m\n"+
				"  Player 2: Human \033[34mB\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[6],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[7],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[8],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[9],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;34mB\033[0m Choose an open space: ",
			actualOutput[10],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;33mY\033[0m Choose an open space: ",
			actualOutput[11],
		)
		assert.Equal(
			t,
			"\n     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;34mB\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;33mY\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[33mY\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[12],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1:  \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 1: ",
			actualOutput[13],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37m\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 1 (e.g. X): ",
			actualOutput[14],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[37mG\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 1: ",
			actualOutput[15],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[32mG\033[0m\n"+
				"  Player 2:  \033[37m\033[0m\n"+
				"  \n\n"+
				"    1) Human\n\n"+
				"  \n"+
				"  Select a player type for Player 2: ",
			actualOutput[16],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[32mG\033[0m\n"+
				"  Player 2: Human \033[37m\033[0m\n"+
				"  \n\n"+
				"  Select a token for Player 2 (e.g. X): ",
			actualOutput[17],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[32mG\033[0m\n"+
				"  Player 2: Human \033[37mM\033[0m\n"+
				"  \n\n"+
				"    1) \033[31mRed\033[0m\n"+
				"    2) \033[32mGreen\033[0m\n"+
				"    3) \033[33mYellow\033[0m\n"+
				"    4) \033[34mBlue\033[0m\n"+
				"    5) \033[35mMagenta\033[0m\n\n"+
				"  \n"+
				"  Select a color for Player 2: ",
			actualOutput[18],
		)
		assert.Equal(
			t,
			"Welcome to Tic-Tac-Go(lang).\n\n"+
				"  Player 1: Human \033[32mG\033[0m\n"+
				"  Player 2: Human \033[35mM\033[0m\n"+
				"  \n\n"+
				"    1) Let's play!\n"+
				"    2) I messed up.\n\n"+
				"  \n"+
				"  All good? ",
			actualOutput[19],
		)
		assert.Equal(
			t,
			"\n     \033[1;37m1\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mG\033[0m Choose an open space: ",
			actualOutput[20],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;37m5\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;35mM\033[0m Choose an open space: ",
			actualOutput[21],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m4\033[0m | \033[1;35mM\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mG\033[0m Choose an open space: ",
			actualOutput[22],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;32mG\033[0m | \033[1;35mM\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;37m7\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;35mM\033[0m Choose an open space: ",
			actualOutput[23],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;37m2\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;32mG\033[0m | \033[1;35mM\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;35mM\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;32mG\033[0m Choose an open space: ",
			actualOutput[24],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;32mG\033[0m | \033[1;37m3\033[0m \n"+
				"    -----------\n"+
				"     \033[1;32mG\033[0m | \033[1;35mM\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;35mM\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"  \033[1;35mM\033[0m Choose an open space: ",
			actualOutput[25],
		)
		assert.Equal(
			t,
			"\n     \033[1;32mG\033[0m | \033[1;32mG\033[0m | \033[1;35mM\033[0m \n"+
				"    -----------\n"+
				"     \033[1;32mG\033[0m | \033[1;35mM\033[0m | \033[1;37m6\033[0m \n"+
				"    -----------\n"+
				"     \033[1;35mM\033[0m | \033[1;37m8\033[0m | \033[1;37m9\033[0m \n\n"+
				"\033[35mM\033[0m Wins!\n\n\n"+
				"    1) Rematch!\n"+
				"    2) New Game\n"+
				"    3) Quit\n\n"+
				"  \n"+
				"  What next? ",
			actualOutput[26],
		)
	})
}
