package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fadet/ai_algorithms/simulatedAnnealing/board"
	"github.com/fadet/ai_algorithms/simulatedAnnealing/plotting"
)

const (
	MAINWINDOW_SIZE float32 = 400
)

func main() {
	a := app.New()
	w := a.NewWindow("N Queens Puzzle")
	w.SetMaster()
	var b, p fyne.Window

	boardSizeInputField := widget.NewEntry()
	initTempInputField := widget.NewEntry()
	finalTempInputField := widget.NewEntry()
	alphaInputField := widget.NewEntry()
	stepsInputField := widget.NewEntry()
	plottingChoice := widget.NewCheck("Plot graph", func(b bool) {})

	inputs := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Board size:", Widget: boardSizeInputField},
			{Text: "Initial Temperature:", Widget: initTempInputField},
			{Text: "Final Temperature:", Widget: finalTempInputField},
			{Text: "Alpha rate:", Widget: alphaInputField},
			{Text: "Steps per change:", Widget: stepsInputField},
		},
		OnSubmit: func() {
			if b != nil {
				b.Close()
			}
			if p != nil {
				p.Close()
			}

			var (
				errorMsg                      string
				boardSize, stepsPerChange     int
				initialTemp, finalTemp, alpha float64
			)

			_, err := fmt.Sscanf(boardSizeInputField.Text, "%d", &boardSize)
			if err != nil {
				errorMsg += "Board size must be an integer.\n"
			}
			_, err = fmt.Sscanf(initTempInputField.Text, "%f", &initialTemp)
			if err != nil {
				errorMsg += "Initial Temperature must be a float.\n"
			}
			_, err = fmt.Sscanf(finalTempInputField.Text, "%f", &finalTemp)
			if err != nil {
				errorMsg += "Final Temperature must be a float.\n"
			}
			_, err = fmt.Sscanf(alphaInputField.Text, "%f", &alpha)
			if err != nil {
				errorMsg += "Alpha rate must be a float.\n"
			}
			_, err = fmt.Sscanf(stepsInputField.Text, "%d", &stepsPerChange)
			if err != nil {
				errorMsg += "Steps per change must be an integer.\n"
			}

			if errorMsg == "" {
				if boardSize < 4 {
					errorMsg = "Board size must be at least 4.\n"
				}
				if initialTemp <= 0 && finalTemp < 0 {
					errorMsg = "Temperature must be positive.\n"
				}
				if initialTemp <= finalTemp {
					errorMsg = "Initial Temperature must be greater than Final Temperature.\n"
				}
				if alpha <= 0 && alpha >= 1 {
					errorMsg = "Alpha must be on an interval (0, 1).)\n"
				}
			}

			if errorMsg != "" {
				e := a.NewWindow("Error")
				e.SetContent(
					widget.NewLabel(errorMsg),
				)
				e.SetFixedSize(true)
				e.Show()
				return
			}

			b = board.GetBoardWindow(a, boardSize, initialTemp, finalTemp, alpha, stepsPerChange)
			b.Show()

			if plottingChoice.Checked {
				p = plotting.GetGraphsWindow(a)
				p.Show()
			}
		},
		SubmitText: "Solve",
	}

	w.SetContent(container.NewVBox(
		inputs,
		plottingChoice,
	),
	)

	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(MAINWINDOW_SIZE, inputs.Size().Height))
	w.ShowAndRun()
}
