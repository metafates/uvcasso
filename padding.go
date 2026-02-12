package uvcasso

import uv "github.com/charmbracelet/ultraviolet"

type Padding struct {
	Top, Right, Bottom, Left int
}

func (p Padding) Apply(area uv.Rectangle) uv.Rectangle {
	horizontal := p.Right + p.Left
	vertical := p.Top + p.Bottom

	if area.Dx() < horizontal || area.Dy() < vertical {
		return uv.Rectangle{}
	}

	return uv.Rect(
		area.Min.X+p.Left,
		area.Min.Y+p.Top,
		max(0, area.Dx()-horizontal),
		max(0, area.Dy()-vertical),
	)
}

func NewPadding(sides ...int) Padding {
	switch len(sides) {
	case 0:
		return Padding{}
	case 1:
		side := sides[0]
		return Padding{Top: side, Right: side, Bottom: side, Left: side}
	case 2:
		vertical := sides[0]
		horizontal := sides[1]

		return Padding{
			Top:    vertical,
			Right:  horizontal,
			Bottom: vertical,
			Left:   horizontal,
		}
	case 4:
		return Padding{
			Top:    sides[0],
			Right:  sides[1],
			Bottom: sides[2],
			Left:   sides[3],
		}
	default:
		panic("unexpected sides count")
	}
}

func (p Padding) WithTop(s int) Padding {
	p.Top = s
	return p
}

func (p Padding) WithRight(s int) Padding {
	p.Right = s
	return p
}

func (p Padding) WithBottom(s int) Padding {
	p.Bottom = s
	return p
}

func (p Padding) WithLeft(s int) Padding {
	p.Bottom = s
	return p
}
