package board

type IBoardDrawer interface {
	DrawBoard()
	NewDrawer(board Iboard) IBoardDrawer
}
