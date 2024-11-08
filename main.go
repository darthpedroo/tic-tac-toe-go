package main

import (
	//"fmt"
	"tictactoe/board"
)

func main() {

	//bot := players.Bot{players.GetNamer{"Bot"}}
	//player := players.Player{players.GetNamer{"Player"}}
	//getNamer := players.GetNamer{"Player"}
	//crossPiece := pieces.NewNormalTicTacToePiece()

	//player := players.NewPlayer("X", getNamer)

	tttboard := board.NewTicTacToeBoard(3, 3)

	consoleBoardDrawer := board.NewDrawer(tttboard)
	//player.PlayPiece(tttboard, 0, 0)
	//player.PlayPiece(tttboard, 2, 0)
	consoleBoardDrawer.DrawBoard()
}
