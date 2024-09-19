package board

/*
#include "emsa.h"
#include <stdlib.h>
*/
import "C"
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"unsafe"
)

const GRID_SIZE float32 = 1000
const PADDING float32 = 10

func GetBoardWindow(a fyne.App, boardSize int, initialTemp, finalTemp, alpha float64, stepsPerChange int) fyne.Window {
	cArray := C.emsa(C.int(boardSize), C.double(initialTemp), C.double(finalTemp), C.double(alpha), C.int(stepsPerChange))
	board := (*[1 << 30]C.int)(unsafe.Pointer(cArray))[:boardSize:boardSize]
	defer C.free(unsafe.Pointer(cArray))

	grid := container.NewGridWithColumns(boardSize)
	size := GRID_SIZE/float32(boardSize) - PADDING*(float32(boardSize)-1)

	for i, j := range board {
		for k := 0; k < boardSize; k++ {
			if k == int(j) {
				var img canvas.Image
				if (i+k)%2 == 1 {
					img = *whiteQueen
				} else {
					img = *blackQueen
				}
				img.Resize(fyne.NewSize(size, size))
				grid.Add(&img)
			} else {
				var colour color.Color
				if (i+k)%2 == 1 {
					colour = color.NRGBA{0, 0, 0, 255}
				} else {
					colour = color.NRGBA{255, 255, 255, 255}
				}
				rect := canvas.NewRectangle(colour)
				rect.Resize(fyne.NewSize(size, size))
				grid.Add(rect)
			}
		}

	}

	w := a.NewWindow("Solution")
	w.SetContent(grid)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(GRID_SIZE, GRID_SIZE))

	return w
}
