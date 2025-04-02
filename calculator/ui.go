package calculator

import (
	"fyne.io/fyne/v2/container"
)

func (c *Calculator) SetupUI() {
	buttons := c.createButtons()
	content := container.NewVbox(
		c.Display,
		container.NewGridWithColumns(1, buttons...),
	)
	c.Window.SetContent(content)
}
