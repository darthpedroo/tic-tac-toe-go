package pieces

type Piece interface {
	GetPosX() int
	GetPosY() int
	GetPiece() string
}