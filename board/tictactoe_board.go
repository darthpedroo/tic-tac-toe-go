// tictactoe/board/board.go
package board

import (
	//"fmt"
	"tictactoe/interfaces"
	"tictactoe/pieces"
)

type TicTacToeBoard struct {
	Columns int
	Rows    int
	board   [][]pieces.Piece
}

func (b TicTacToeBoard) GetColumns() int {
	return b.Columns
}

func NewTicTacToeBoard(columns int, rows int) *TicTacToeBoard {

	board := make([][]pieces.Piece, columns)

	for i := range board {
		board[i] = make([]pieces.Piece, rows)
	}

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			board[i][j] = pieces.NewNormalTicTacToePiece(i, j, "ñ")
		}
	}

	return &TicTacToeBoard{
		Columns: columns,
		Rows:    rows,
		board:   board,
	}

}

func (b TicTacToeBoard) CreateIterator() interfaces.Iterator {

	return &interfaces.BoardIterator{

		Columns: 0,
		Rows:    0,
		Board:   b.board,
	}

}

func (b *TicTacToeBoard) AddPiece(piece pieces.Piece) {

	posX := piece.GetPosX()
	posY := piece.GetPosY()

	b.board[posX][posY] = piece

}
