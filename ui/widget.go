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
)

type Widget interface {
	HandleInput(region image.Rectangle) Widget
	Update(focused Widget) error
	Draw(screen *ebiten.Image, region image.Rectangle)
}

type Panel struct {
	Children        []Widget
	BackgroundColor color.Color
}

func (p *Panel) HandleInput(region image.Rectangle) Widget {
	for _, c := range p.Children {
		if w := c.HandleInput(region); w != nil {
			return w
		}
	}
	return nil
}

func (p *Panel) Update(focused Widget) error {
	for _, c := range p.Children {
		if err := c.Update(focused); err != nil {
			return err
		}
	}
	return nil
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

type Label struct {
	Region          image.Rectangle
	Text            string
	HorizontalAlign HorizontalAlign
	VerticalAlign   VerticalAlign
}

func (l *Label) HandleInput(region image.Rectangle) Widget {
	return nil
}

func (l *Label) Update(focused Widget) error {
	return nil
}

func (l *Label) Draw(screen *ebiten.Image, region image.Rectangle) {
	r := absRegion(l.Region, region)

	x, y := textAt(l.Text, r, l.HorizontalAlign, l.VerticalAlign)
	text.Draw(screen, l.Text, bitmapfont.Gothic12r, x, y, color.Black)
}

type Button struct {
	Region image.Rectangle
	Text   string

	OnClick func(b *Button)

	pressed bool
}

func absRegion(rel, region image.Rectangle) image.Rectangle {
	x, y := region.Min.X+rel.Min.X, region.Min.Y+rel.Min.Y
	return image.Rect(x, y, x+rel.Dx(), y+rel.Dy())
}

func (b *Button) HandleInput(region image.Rectangle) Widget {
	r := absRegion(b.Region, region)
	if b.pressed {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			return b
		}
		b.pressed = false
		if !image.Pt(ebiten.CursorPosition()).In(r) {
			return nil
		}
		if b.OnClick != nil {
			b.OnClick(b)
		}
		return nil
	}
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	if !image.Pt(ebiten.CursorPosition()).In(r) {
		return nil
	}
	b.pressed = true
	return b
}

func (b *Button) Update(focused Widget) error {
	return nil
}

func (b *Button) Draw(screen *ebiten.Image, region image.Rectangle) {
	r := absRegion(b.Region, region)
	drawNinePatch(screen, tmpButtonImage, r)

	x, y := textAt(b.Text, r, Center, Middle)
	text.Draw(screen, b.Text, bitmapfont.Gothic12r, x, y, color.Black)
}

const textBoxPadding = 8

type TextBox struct {
	Region image.Rectangle
	Value  string

	index   int
	focused bool
	tick    int
}

func (t *TextBox) HandleInput(region image.Rectangle) Widget {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	r := absRegion(t.Region, region)
	cx, cy := ebiten.CursorPosition()
	if !image.Pt(cx, cy).In(r) {
		return nil
	}

	t.index = closestTextIndex(t.Value, cx-(r.Min.X+textBoxPadding))
	return t
}

func (t *TextBox) Update(focused Widget) error {
	if t != focused {
		t.focused = false
		t.tick = 0
		return nil
	}
	t.focused = true
	t.tick++
	t.tick = t.tick % 60

	v := []rune(t.Value)
	rs := ebiten.InputChars()
	t.Value = string(v[:t.index]) + string(rs) + string(v[t.index:])
	t.index += len(rs)
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && t.index > 0 {
		if rs := []rune(t.Value); len(rs) >= t.index {
			t.Value = string(rs[:t.index-1]) + string(rs[t.index:])
			t.index--
		}
	}
	// TODO: Emulate repeating
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && t.index > 0 {
		t.index--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && t.index < len([]rune(t.Value)) {
		t.index++
	}
	return nil
}

func (t *TextBox) Draw(screen *ebiten.Image, region image.Rectangle) {
	r := absRegion(t.Region, region)
	drawNinePatch(screen, tmpTextBoxImage, r)

	x, y := textAt(t.Value, r, Left, Middle)
	x += textBoxPadding
	text.Draw(screen, t.Value, bitmapfont.Gothic12r, x, y, color.Black)

	if !t.focused {
		return
	}
	if t.tick >= 30 {
		return
	}
	dx, _ := textSize(string([]rune(t.Value)[:t.index]))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.Min.X+textBoxPadding+dx), float64(r.Min.Y))
	screen.DrawImage(tmpIBeamImage, op)
}

var (
	tmpButtonImage  *ebiten.Image
	tmpTextBoxImage *ebiten.Image
	tmpIBeamImage   *ebiten.Image
)

func init() {
	tmpButtonImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
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

func init() {
	tmpTextBoxImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	pix := make([]byte, 4*16*16)
	idx := 0
	for j := 0; j < 16; j++ {
		for i := 0; i < 16; i++ {
			if j == 15 {
				pix[idx] = 0x33
				pix[idx+1] = 0x33
				pix[idx+2] = 0x33
				pix[idx+3] = 0xff
			} else {
				pix[idx] = 0xee
				pix[idx+1] = 0xee
				pix[idx+2] = 0xee
				pix[idx+3] = 0xff
			}
			idx += 4
		}
	}
	tmpTextBoxImage.ReplacePixels(pix)
}

func init() {
	tmpIBeamImage, _ = ebiten.NewImage(1, 14, ebiten.FilterDefault)
	tmpIBeamImage.Fill(color.Black)
}
