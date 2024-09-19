package plotting

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"sync"
)

const IMAGE_SIZE = float32(6 * vg.Inch)

func saveData() {
	sts := getStats()
	t, e, c := getPoints(sts)

	wg := sync.WaitGroup{}

	f := func(xl, yl, fn string, pts *plotter.XYs, colour color.Color) {
		defer wg.Done()

		p := plot.New()
		p.X.Label.Text = xl
		p.Y.Label.Text = yl
		l, lp, _ := plotter.NewLinePoints(pts)
		l.LineStyle.Color = colour
		lp.GlyphStyle.Shape = draw.CrossGlyph{}
		lp.GlyphStyle.Color = color.Gray{20}

		p.Add(l, lp)
		p.Save(6*vg.Inch, 6*vg.Inch, fn)
	}

	wg.Add(3)
	go f("Iterations", "Temperature", "temperature.png", t, color.NRGBA{255, 0, 0, 255})
	go f("Iterations", "Best Energy", "energy.png", e, color.NRGBA{0, 255, 0, 255})
	go f("Iterations", "Bad Choices", "choices.png", c, color.NRGBA{0, 0, 255, 255})

	wg.Wait()
}

func GetGraphsWindow(a fyne.App) fyne.Window {
	saveData()

	var t, e, c *canvas.Image
	t = canvas.NewImageFromFile("./temperature.png")
	e = canvas.NewImageFromFile("./energy.png")
	c = canvas.NewImageFromFile("./choices.png")

	w := a.NewWindow("Graphs")
	w.SetContent(
		container.NewGridWithColumns(3, t, e, c),
	)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(IMAGE_SIZE*3, IMAGE_SIZE))

	return w
}
