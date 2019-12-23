// SPDX-License-Identifier: Apache-2.0

package ui

import (
	"image"

	"github.com/hajimehoshi/bitmapfont"
	"golang.org/x/image/font"
)

type HorizontalAlign int

const (
	Left HorizontalAlign = iota
	Center
	Right
)

type VerticalAlign int

const (
	Top VerticalAlign = iota
	Middle
	Bottom
)

func textAt(text string, region image.Rectangle, h HorizontalAlign, v VerticalAlign) (int, int) {
	f := bitmapfont.Gothic12r
	bound, _ := font.BoundString(f, text)
	bw := (bound.Max.X - bound.Min.X).Round()
	bh := (bound.Max.Y - bound.Min.Y).Round()
	x := region.Min.X + 4
	switch h {
	case Left:
	case Center:
		x += (region.Dx() - bw) / 2
	case Right:
		x += region.Dx() - bw
	}
	y := region.Min.Y + 12
	switch v {
	case Top:
	case Middle:
		y += (region.Dy() - bh) / 2
	case Bottom:
		y += region.Dy() - bh
	}
	return x, y
}
