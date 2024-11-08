package players

import (
	"tictactoe/board"
	"tictactoe/pieces"
)

type Player struct {
	pieceStr string
	GetNamer
}

func (player Player) PlayPiece(b board.Iboard, posX int, posY int) {

	newPiece := pieces.NewNormalTicTacToePiece(posX,posY, player.pieceStr)

	b.AddPiece(newPiece)

}

func NewPlayer(piece string, getNamer GetNamer) Player {
	return Player{
		pieceStr:    piece,
		GetNamer: getNamer,
	}
}
