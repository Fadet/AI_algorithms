package board

import "fyne.io/fyne/v2/canvas"

var whiteQueen, blackQueen *canvas.Image

func init() {
	whiteQueen = canvas.NewImageFromFile("./board/white_queen.png")
	blackQueen = canvas.NewImageFromFile("./board/black_queen.png")
}
