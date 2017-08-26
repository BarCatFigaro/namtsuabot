package namtsuabot

import (
	"fmt"
	"strings"
)

/*
		board := `
					8 ║♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
					7 ║♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
					6 ║
					5 ║
					4 ║
					3 ║
					2 ║♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
					1 ║♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖
	                 —╚═══════════════
					—— a  b c  d e  f g  h`
*/

// bX reprents the black pieces, wX represents the white pieces
const (
	bRook   = 9820
	bKnight = 9822
	bBishop = 9821
	bQueen  = 9819
	bKing   = 9818
	bPawn   = 9823

	wRook   = 9814
	wKnight = 9816
	wBishop = 9815
	wQueen  = 9813
	wKing   = 9812
	wPawn   = 9817
)

// board represents the board state in a 2D rune slice
var board = [][]rune{
	{9820, 32, 9822, 32, 9821, 32, 9819, 32, 9818, 32, 9821, 32, 9822, 32, 9820},
	{9823, 32, 9823, 32, 9823, 32, 9823, 32, 9823, 32, 9823, 32, 9823, 32, 9823},
	{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
	{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
	{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
	{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
	{9817, 32, 9817, 32, 9817, 32, 9817, 32, 9817, 32, 9817, 32, 9817, 32, 9817},
	{9814, 32, 9816, 32, 9815, 32, 9813, 32, 9812, 32, 9815, 32, 9816, 32, 9814},
}

// players represent the players in the current match mapped to a colour (true = white, false = black)
var players = make(map[string]bool)

// chessCommand is the chess game delegation function
func chessCommand(user, arg string) string {
	if arg == "start" {
		return buildBoard(user)
	}
	return "func call: chessCommand"
}

// buildBoard constructs the intial board state and registers the users (2)
func buildBoard(user string) string {
	out := []string{}
	for i := 0; i < len(board); i++ {
		out = append(out, string(board[i]))
	}

	board := fmt.Sprintf("%s started a chess match:\n\n", user) +
		strings.Join(out, "\n")
	return board
}

// blackMove modifies the board array with respect to the black side
func blackMove() string {
	return ""
}

/*
func buildBoard() string {
	board := "Chess board" + `

		♜ ♞ ♝ ♛ ♚ ♝ ♜ ♞
		♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟




		♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
		♖ ♘ ♗ ♕ ♔ ♖ ♘ ♗`

	return board
}
*/
