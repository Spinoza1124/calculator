package calculator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	Equation string
	Display  *widget.Label
	Window   fyne.Window
}

func NewCalculator() *Calculator {
	app := app.New()
	window := app.NewWindow("Go 计算器")
	window.Resize(fyne.NewSize(200, 200))

	display := widget.NewLabel("0")
	display.Alignment = fyne.TextAlignTrailing
	display.TextStyle = fyne.TextStyle{Monospace: true}

	calc := &Calculator{
		Display: display,
		Window:  window,
	}

	calc.SetupUI()
	return calc
}
