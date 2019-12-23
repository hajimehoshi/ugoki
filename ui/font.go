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

func textSize(text string) (width, height int) {
	f := bitmapfont.Gothic12r
	b, _ := font.BoundString(f, text)
	return (b.Max.X - b.Min.X).Round(), (b.Max.Y - b.Min.Y).Round()
}

func textAt(text string, region image.Rectangle, h HorizontalAlign, v VerticalAlign) (int, int) {
	bw, bh := textSize(text)
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

func closestTextIndex(str string, x int) int {
	rs := []rune(str)
	pw := -1
	// TODO: Use more efficient algorithm.
	for i := range rs {
		w, _ := textSize(string(rs[:i]))
		if x < w {
			if w-x > x-pw {
				return i - 1
			} else {
				return i
			}
		}
		pw = w
	}
	return len(rs)
}
