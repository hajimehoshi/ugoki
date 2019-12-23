// SPDX-License-Identifier: Apache-2.0

package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/bitmapfont"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type Widget interface {
	HandleInput(region image.Rectangle) bool
	Draw(screen *ebiten.Image, region image.Rectangle)
}

type Panel struct {
	Children        []Widget
	BackgroundColor color.Color
}

func (p *Panel) HandleInput(region image.Rectangle) bool {
	for _, c := range p.Children {
		if c.HandleInput(region) {
			return true
		}
	}
	return true
}

func (p *Panel) Draw(screen *ebiten.Image, region image.Rectangle) {
	if region.Dx() == 0 || region.Dy() == 0 {
		return
	}

	if p.BackgroundColor != nil {
		x := float64(region.Min.X)
		y := float64(region.Min.Y)
		w := float64(region.Dx())
		h := float64(region.Dy())
		ebitenutil.DrawRect(screen, x, y, w, h, p.BackgroundColor)
	}

	for _, c := range p.Children {
		c.Draw(screen, region)
	}
}

type Button struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string

	OnClick func(b *Button)

	pressed bool
}

func (b *Button) absRegion(region image.Rectangle) image.Rectangle {
	x, y := region.Min.X+b.X, region.Min.Y+b.Y
	return image.Rect(x, y, x+b.Width, y+b.Height)
}

func (b *Button) HandleInput(region image.Rectangle) bool {
	r := b.absRegion(region)
	if !image.Pt(ebiten.CursorPosition()).In(r) {
		b.pressed = false
		return false
	}

	if b.pressed {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			return true
		}
		if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			return true
		}
		if b.OnClick != nil {
			b.OnClick(b)
		}
		b.pressed = false
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.pressed = true
	}
	return true
}

func (b *Button) Draw(screen *ebiten.Image, region image.Rectangle) {
	r := b.absRegion(region)
	drawNinePatch(screen, tmpButtonImage, r)
	f := bitmapfont.Gothic12r
	bound, _ := font.BoundString(f, b.Text)
	bw := (bound.Max.X - bound.Min.X).Round()
	bh := (bound.Max.Y - bound.Min.Y).Round()
	x := r.Min.X + 4 + (r.Dx()-bw)/2
	y := r.Min.Y + 12 + (r.Dy()-bh)/2
	text.Draw(screen, b.Text, f, x, y, color.Black)
}

var tmpButtonImage *ebiten.Image

func init() {
	tmpButtonImage, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	pix := make([]byte, 4*16*16)
	idx := 0
	for j := 0; j < 16; j++ {
		for i := 0; i < 16; i++ {
			if i == 0 || i == 15 || j == 0 || j == 15 {
				pix[idx] = 0x33
				pix[idx+1] = 0x33
				pix[idx+2] = 0x33
				pix[idx+3] = 0xff
			} else {
				pix[idx] = 0xcc
				pix[idx+1] = 0xcc
				pix[idx+2] = 0xcc
				pix[idx+3] = 0xff
			}
			idx += 4
		}
	}
	tmpButtonImage.ReplacePixels(pix)
}
