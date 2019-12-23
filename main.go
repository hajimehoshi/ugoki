// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ugoki/ui"
)

func main() {
	toolbar := &ui.HSplitter{
		Children: []ui.Widget{
			&ui.Panel{
				Children: []ui.Widget{
					&ui.Button{
						Region: image.Rect(8, 8, 48, 32),
						Text:   "+",
						OnClick: func(b *ui.Button) {
							fmt.Println("+")
						},
					},
					&ui.Button{
						Region: image.Rect(48, 8, 88, 32),
						Text:   "=",
						OnClick: func(b *ui.Button) {
							fmt.Println("=")
						},
					},
					&ui.Button{
						Region: image.Rect(88, 8, 128, 32),
						Text:   "-",
						OnClick: func(b *ui.Button) {
							fmt.Println("-")
						},
					},
					&ui.Button{
						Region: image.Rect(136, 8, 176, 32),
						Text:   "Left",
					},
					&ui.Button{
						Region: image.Rect(176, 8, 216, 32),
						Text:   "Down",
					},
					&ui.Button{
						Region: image.Rect(216, 8, 256, 32),
						Text:   "Right",
					},
				},
			},
			nil,
			&ui.Panel{
				Children: []ui.Widget{
					&ui.Button{
						Region: image.Rect(8, 8, 48, 32),
						Text:   "A",
					},
					&ui.Button{
						Region: image.Rect(48, 8, 88, 32),
						Text:   "B",
					},
					&ui.Button{
						Region: image.Rect(88, 8, 128, 32),
						Text:   "C",
					},
					&ui.Button{
						Region: image.Rect(136, 8, 176, 32),
						Text:   "D",
					},
					&ui.Button{
						Region: image.Rect(176, 8, 216, 32),
						Text:   "E",
					},
					&ui.Button{
						Region: image.Rect(216, 8, 256, 32),
						Text:   "F",
					},
				},
			},
		},
		Widths: []int{264, -1, 264},
	}

	inspector := &ui.Panel{
		Children: []ui.Widget{
			&ui.Label{
				Text: "Foo",
			},
		},
	}

	mainPanel := &ui.HSplitter{
		Children: []ui.Widget{
			&ui.Panel{
				BackgroundColor: color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
			},
			&ui.VSplitter{
				Children: []ui.Widget{
					&ui.Panel{
						BackgroundColor: color.RGBA{0x33, 0x33, 0x33, 0xff},
					},
					&ui.Panel{
						BackgroundColor: color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
					},
				},
				Heights: []int{-1, 160},
			},
			inspector,
		},
		Widths: []int{200, -1, 200},
	}

	statusBar := &ui.Panel{
		BackgroundColor: color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	}

	u := ui.New(&ui.Panel{
		Children: []ui.Widget{
			&ui.VSplitter{
				Children: []ui.Widget{
					toolbar,
					mainPanel,
					statusBar,
				},
				Heights: []int{40, -1, 24},
			},
		},
	})
	u.SetTitle("Ugoki")
	if err := u.Main(); err != nil {
		panic(err)
	}
}
