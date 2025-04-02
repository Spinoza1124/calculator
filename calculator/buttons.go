package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (c *Calculator) createButtons() []fyne.CanvasObject {
	buttons := []struct {
		label    string
		onTapped func()
	}{
		{"C", c.Clear},
		{"√", c.Sqrt},
		{"^", c.Power},
		{"%", c.Modulo},
		{"7", func() { c.AppendNumber("7") }},
		{"8", func() { c.AppendNumber("8") }},
		{"9", func() { c.AppendNumber("9") }},
		{"/", func() { c.AppendOperator("/") }},
		{"4", func() { c.AppendNumber("4") }},
		{"5", func() { c.AppendNumber("5") }},
		{"6", func() { c.AppendNumber("6") }},
		{"*", func() { c.AppendOperator("*") }},
		{"1", func() { c.AppendNumber("1") }},
		{"2", func() { c.AppendNumber("2") }},
		{"3", func() { c.AppendNumber("3") }},
		{"-", func() { c.AppendOperator("-") }},
		{"0", func() { c.AppendNumber("0") }},
		{".", c.AppendDecimal},
		{"=", c.Calculate},
		{"+", func() { c.AppendOperator("+") }},
	}

	var buttonWidgets []fyne.CanvasObject
	for _, btn := range buttons {
		button := widget.NewButton(btn.label, btn.onTapped)
		buttonWidgets = append(buttonWidgets, button)
	}

	return buttonWidgets
}

func (c *Calculator) AppendNumber(num string) {
	if c.Display.Text == "0" || c.Display.Text == "错误" {
		c.Display.SetText(num)
	} else {
		c.Display.SetText(c.Display.Text + num)
	}
}

func (c *Calculator) AppendDecimal() {
	if !strings.Contains(c.Display.Text, ".") {
		c.Display.SetText(c.Display.Text + ".")
	}
}

func (c *Calculator) AppendOperator(op string) {
	text := c.Display.Text
	if text == "错误" {
		return
	}

	if len(text) > 0 {
		lastChar := text[len(text)-1:]
		if strings.ContainsAny(lastChar, "+-*/%^") {
			text = text[:len(text)-1]
		}
	}

	c.Display.SetText(text + op)
}

func (c *Calculator) Clear() {
	c.Display.SetText("0")
	c.Equation = ""
}

func (c *Calculator) Sqrt() {
	text := c.Display.Text
	if text == "错误" {
		return
	}

	num, err := strconv.ParseFloat(text, 64)
	if err != nil {
		c.Display.SetText("错误")
		return
	}

	if num < 0 {
		c.Display.SetText("错误")
		return
	}

	result := math.Sqrt(num)
	c.Display.SetText(fmt.Sprintf("%g", result))
}

func (c *Calculator) Power() {
	c.AppendOperator("^")
}

func (c *Calculator) Modulo() {
	c.AppendOperator("%")
}
