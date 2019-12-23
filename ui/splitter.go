// SPDX-License-Identifier: Apache-2.0

package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type HSplitter struct {
	Children []Widget
	Widths   []int
}

type VSplitter struct {
	Children []Widget
	Heights  []int
}

func calcParts(total int, parts []int) []int {
	ps := make([]int, len(parts))
	copy(ps, parts)

	s := 0
	n := 0
	for _, p := range parts {
		if p == -1 {
			n++
		} else {
			s += p
		}
	}
	if n == 0 {
		return ps
	}

	p := (total - s) / n
	if p < 0 {
		p = 0
	}
	for i, v := range parts {
		if v == -1 {
			ps[i] = p
		}
	}
	return ps
}

func (h *HSplitter) childRegion(region image.Rectangle, index int) image.Rectangle {
	ws := calcParts(region.Dx(), h.Widths)
	xs := make([]int, len(h.Widths))
	for i := range ws {
		if i > 0 {
			xs[i] = xs[i-1] + ws[i-1]
		}
	}
	return image.Rect(region.Min.X+xs[index], region.Min.Y, region.Min.X+xs[index]+ws[index], region.Max.Y)
}

func (h *HSplitter) HandleInput(region image.Rectangle) Widget {
	for i, c := range h.Children {
		if c == nil {
			continue
		}
		if w := c.HandleInput(h.childRegion(region, i)); w != nil {
			return w
		}
	}
	return nil
}

func (h *HSplitter) Update(focused Widget) error {
	for _, c := range h.Children {
		if c == nil {
			continue
		}
		if err := c.Update(focused); err != nil {
			return err
		}
	}
	return nil
}

func (h *HSplitter) Draw(screen *ebiten.Image, region image.Rectangle) {
	for i, c := range h.Children {
		if c == nil {
			continue
		}
		c.Draw(screen, h.childRegion(region, i))
	}
}

func (v *VSplitter) childRegion(region image.Rectangle, index int) image.Rectangle {
	hs := calcParts(region.Dy(), v.Heights)
	ys := make([]int, len(v.Heights))
	for i := range hs {
		if i > 0 {
			ys[i] = ys[i-1] + hs[i-1]
		}
	}
	return image.Rect(region.Min.X, region.Min.Y+ys[index], region.Max.X, region.Min.Y+ys[index]+hs[index])
}

func (v *VSplitter) HandleInput(region image.Rectangle) Widget {
	for i, c := range v.Children {
		if c == nil {
			continue
		}
		if w := c.HandleInput(v.childRegion(region, i)); w != nil {
			return w
		}
	}
	return nil
}

func (v *VSplitter) Update(focused Widget) error {
	for _, c := range v.Children {
		if c == nil {
			continue
		}
		if err := c.Update(focused); err != nil {
			return err
		}
	}
	return nil
}

func (v *VSplitter) Draw(screen *ebiten.Image, region image.Rectangle) {
	for i, c := range v.Children {
		if c == nil {
			continue
		}
		c.Draw(screen, v.childRegion(region, i))
	}
}
