package pieces

type NormalTicTacToePiece struct {
	posX  int
	posY  int
	piece string
}

func (p NormalTicTacToePiece) GetPosX() int {
	return p.posX
}

func (p NormalTicTacToePiece) GetPosY() int {
	return p.posY
}

func (p NormalTicTacToePiece) GetPiece() string {
	return p.piece
}

func NewNormalTicTacToePiece(posX int, posY int, piece string) NormalTicTacToePiece {

	return NormalTicTacToePiece{
		posX:  posX,
		posY:  posX,
		piece: piece,
	}

}
