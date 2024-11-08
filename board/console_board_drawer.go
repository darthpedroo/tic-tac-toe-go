package board

import (
	"fmt"
)

type ConsoleBoardDrawer struct {
	Board Iboard
}

func (drawer ConsoleBoardDrawer) DrawBoard() {

	boardIterator := drawer.Board.CreateIterator()

	columns := drawer.Board.GetColumns()
	currentColumn := 0

	for boardIterator.HasNext() {

		piece := boardIterator.GetNext()		
		fmt.Print(" | ")
		fmt.Print(piece.GetPiece())
		fmt.Print(" | ")
		currentColumn += 1
		drawSpacesBetweenTiles(currentColumn, columns)
				
	}
}

func drawSpacesBetweenTiles(currentColumn int, columns int) {

	if currentColumn%columns == 0 {
		fmt.Println("")
		for column := 0; column < columns; column++{
			fmt.Printf("  ")
		}
		fmt.Println("")
		
	}

}


func NewDrawer(b Iboard) ConsoleBoardDrawer {
	return ConsoleBoardDrawer{
		Board: b,
	}
}
