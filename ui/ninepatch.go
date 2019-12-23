// SPDX-License-Identifier: Apache-2.0

package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

func drawNinePatch(dst, src *ebiten.Image, region image.Rectangle) {
	sw, sh := src.Size()

	sx := []int{0, sw / 4, sw * 3 / 4, sw}
	sy := []int{0, sh / 4, sh * 3 / 4, sh}
	dox := float64(region.Min.X)
	doy := float64(region.Min.Y)
	dx := []float64{dox, dox+float64(sw / 4), dox+float64(region.Dx() - sw/4)}
	dy := []float64{doy, doy+float64(sh / 4), doy+float64(region.Dy() - sh/4)}
	dw := []float64{1.0, float64(region.Dx()-sw/2) / float64(sw/2), 1.0}
	dh := []float64{1.0, float64(region.Dy()-sh/2) / float64(sh/2), 1.0}

	op := &ebiten.DrawImageOptions{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			op.GeoM.Reset()
			op.GeoM.Scale(dw[i], dh[j])
			op.GeoM.Translate(dx[i], dy[j])
			dst.DrawImage(src.SubImage(image.Rect(sx[i], sy[j], sx[i+1], sy[j+1])).(*ebiten.Image), op)
		}
	}
}
