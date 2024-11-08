package interfaces

import (
	"tictactoe/pieces"
)

type Iterator interface {
	HasNext() bool
	GetNext() pieces.Piece // Returns the current coordinate and the piece at that coordinate
}

type BoardIterator struct {
	Columns int
	Rows    int
	Board   [][]pieces.Piece
}

func (it *BoardIterator) HasNext() bool {

	if it.Rows < len(it.Board) {
		if it.Columns < len(it.Board[it.Rows]) {
			return true
		}
	}
	return false

}

func (it *BoardIterator) GetNext() pieces.Piece {
	if it.HasNext() {

		piece := it.Board[it.Columns][it.Rows]
		it.Columns++

		// Si la columna actual está al final, pasa a la siguiente fila
		if it.Columns >= len(it.Board[it.Rows]) {
			it.Columns = 0
			it.Rows++
		}
		return piece
	}
	return nil
}
