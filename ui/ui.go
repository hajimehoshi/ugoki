// SPDX-License-Identifier: Apache-2.0

package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type game struct {
	panel *Panel
}

func (g *game) Update(screen *ebiten.Image) error {
	if err := g.update(screen.Size()); err != nil {
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.draw(screen)
	return nil
}

func (g *game) update(width, height int) error {
	g.panel.HandleInput(image.Rect(0, 0, width, height))
	return nil
}

func (g *game) draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	w, h := screen.Size()
	g.panel.Draw(screen, image.Rect(0, 0, w, h))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type UI struct {
	game *game
}

func New(panel *Panel) *UI {
	return &UI{
		game: &game{
			panel: panel,
		},
	}
}

func (u *UI) SetWindowSize(width, height int) {
	ebiten.SetWindowSize(width, height)
}

func (u *UI) SetTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (u *UI) Main() error {
	ebiten.SetWindowResizable(true)
	return ebiten.RunGame(u.game)
}
