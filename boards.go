package main

import (
	"github.com/nsf/termbox-go"
)

func init() {
	boards = []Boards{

		// 10 x 20 blank
		Boards{
			colors: [][]termbox.Attribute{
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
				[]termbox.Attribute{blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor,
					blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor, blankColor},
			},
			rotation: [][]int{
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			}},

	}
}
