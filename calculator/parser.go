package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Calculator) Calculate() {
	expression := c.Display.Text
	if expression == "错误" {
		return
	}

	expression = strings.ReplaceAll(expression, "^", "**")
	result, err := evaluateExpression(expression)
	if err != nil {
		c.Display.SetText("错误")
		return
	}

	c.Display.SetText(fmt.Sprintf("%g", result))
}

// 表达式解析器实现
type exprParser struct {
	expr string
	pos  int
}

func evaluateExpression(expr string) (float64, error) {
	parser := &exprParser{expr: expr, pos: 0}
	return parser.parse()
}

func (p *exprParser) parse() (float64, error) {
	return p.parseAddSub()
}

func (p *exprParser) parseAddSub() (float64, error) {
	left, err := p.parseMulDiv()
	if err != nil {
		return 0, err
	}

	for {
		if p.pos >= len(p.expr) {
			return left, nil
		}

		op := p.expr[p.pos]
		if op != '+' && op != '-' {
			return left, nil
		}
		p.pos++

		right, err := p.parseMulDiv()
		if err != nil {
			return 0, err
		}

		switch op {
		case '+':
			left += right
		case '-':
			left -= right
		}
	}
}

func (p *exprParser) parseMulDiv() (float64, error) {
	left, err := p.parsePower()
	if err != nil {
		return 0, err
	}

	for {
		if p.pos >= len(p.expr) {
			return left, nil
		}

		op := p.expr[p.pos]
		if op != '*' && op != '/' && op != '%' {
			return left, nil
		}
		p.pos++

		right, err := p.parsePower()
		if err != nil {
			return 0, err
		}

		switch op {
		case '*':
			left *= right
		case '/':
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			left /= right
		case '%':
			left = math.Mod(left, right)
		}
	}
}

func (p *exprParser) parsePower() (float64, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}

	for {
		if p.pos >= len(p.expr) {
			return left, nil
		}

		if p.pos+1 < len(p.expr) && p.expr[p.pos] == '*' && p.expr[p.pos+1] == '*' {
			p.pos += 2

			right, err := p.parseUnary()
			if err != nil {
				return 0, err
			}

			left = math.Pow(left, right)
		} else {
			return left, nil
		}
	}
}

func (p *exprParser) parseUnary() (float64, error) {
	if p.pos >= len(p.expr) {
		return 0, fmt.Errorf("unexpected end of expression")
	}

	if p.expr[p.pos] == '+' {
		p.pos++
		return p.parsePrimary()
	} else if p.expr[p.pos] == '-' {
		p.pos++
		val, err := p.parsePrimary()
		return -val, err
	}

	return p.parsePrimary()
}

func (p *exprParser) parsePrimary() (float64, error) {
	if p.pos >= len(p.expr) {
		return 0, fmt.Errorf("unexpected end of expression")
	}

	if p.expr[p.pos] == '(' {
		p.pos++
		val, err := p.parseAddSub()
		if err != nil {
			return 0, err
		}
		if p.pos >= len(p.expr) || p.expr[p.pos] != ')' {
			return 0, fmt.Errorf("missing closing parenthesis")
		}
		p.pos++
		return val, nil
	}

	start := p.pos
	for p.pos < len(p.expr) && (isDigit(p.expr[p.pos]) || p.expr[p.pos] == '.') {
		p.pos++
	}

	if start == p.pos {
		return 0, fmt.Errorf("expected number")
	}

	numStr := p.expr[start:p.pos]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
