// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ugoki/ui"
)

func main() {
	toolbar := &ui.HSplitter{
		Children: []ui.Widget{
			&ui.Panel{
				Children: []ui.Widget{
					&ui.Button{
						X:      8,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "+",
						OnClick: func(b *ui.Button) {
							fmt.Println("+")
						},
					},
					&ui.Button{
						X:      48,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "=",
						OnClick: func(b *ui.Button) {
							fmt.Println("=")
						},
					},
					&ui.Button{
						X:      88,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "-",
						OnClick: func(b *ui.Button) {
							fmt.Println("-")
						},
					},
					&ui.Button{
						X:      136,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "Left",
					},
					&ui.Button{
						X:      176,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "Down",
					},
					&ui.Button{
						X:      216,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "Right",
					},
				},
			},
			nil,
			&ui.Panel{
				Children: []ui.Widget{
					&ui.Button{
						X:      8,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "A",
					},
					&ui.Button{
						X:      48,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "B",
					},
					&ui.Button{
						X:      88,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "C",
					},
					&ui.Button{
						X:      136,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "D",
					},
					&ui.Button{
						X:      176,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "E",
					},
					&ui.Button{
						X:      216,
						Y:      8,
						Width:  40,
						Height: 24,
						Text:   "F",
					},
				},
			},
		},
		Widths: []int{264, -1, 264},
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
			&ui.Panel{
				BackgroundColor: color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
			},
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
