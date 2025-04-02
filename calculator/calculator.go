package calculator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/sys/windows"
)

type Calculator struct {
	Equation string
	Display  *widget.Label
	windows  fyne.Window
}

func NewCalculator() *Calculator {
	app := app.New()
	window := app.NewWindow("Go 计算器")
	window.Resize(fyne.NewSize(300, 400))

	display := widget.NewLabel("0")
	display.Alignment = fyne.TextAlignTrailing
	display.TextStyle = fyne.TextStyle{Monspace: true}

	return &Calculator{
		Display: display,
		windows: window,
	}
}
