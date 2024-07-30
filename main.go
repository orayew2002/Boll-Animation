package main

import (
	"fmt"
	"time"
)

const (
	Width  = 26
	Height = 5
)

var (
	x, y   = 0, 0
	vx, vy = 1, 1
)

func main() {
	clearScreen()

	bufBoard := make([]rune, 0, Width*Height)

	board := make([][]bool, Width)
	for i := 0; i < Width; i++ {
		board[i] = make([]bool, Height)
	}

	var cell rune

	for {
		moveCursorToTopLeft()

		board[x][y] = true
		bufBoard = bufBoard[:0]

		for y := range board[0] {
			for x := range board {
				cell = ' '

				if board[x][y] {
					cell = '⚽'
				}

				bufBoard = append(bufBoard, cell)
			}

			bufBoard = append(bufBoard, '\n')
		}

		fmt.Println(string(bufBoard))

		board[x][y] = false

		if y == Height-1 {
			vy = -1
		}

		if y == 0 {
			vy = 1
		}

		if x == Width-1 {
			vx = -1
		}

		if x == 0 {
			vx = 1
		}

		x += vx
		y += vy

		time.Sleep(time.Second / 20)
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func moveCursorToTopLeft() {
	fmt.Print("\033[H")
}
