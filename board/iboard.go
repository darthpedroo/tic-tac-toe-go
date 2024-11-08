package board
import ("tictactoe/interfaces"
		"tictactoe/pieces")

type Iboard interface {
	CreateIterator() interfaces.Iterator
	GetColumns() int
	AddPiece(piece pieces.Piece)
}

