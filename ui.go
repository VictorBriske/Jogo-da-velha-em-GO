package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawBoard(screen *ebiten.Image, board [3][3]int) {
	cellSize := 100
	lineColor := color.White
	lineWidth := float32(2)

	// Linhas do tabuleiro
	for i := 1; i < 3; i++ {
		x := float32(i * cellSize)
		vector.StrokeLine(screen, x, 0, x, 300, lineWidth, lineColor, false)
		vector.StrokeLine(screen, 0, x, 300, x, lineWidth, lineColor, false)
	}

	// Peças
	xColor := color.RGBA{255, 100, 100, 255}
	oColor := color.RGBA{100, 200, 255, 255}
	symbolWidth := float32(5)

	for i := range board {
		for j := range board[i] {
			cx := float32(j*cellSize + cellSize/2)
			cy := float32(i*cellSize + cellSize/2)

			switch board[i][j] {
			case Player:
				offset := float32(30)
				vector.StrokeLine(screen, cx-offset, cy-offset, cx+offset, cy+offset, symbolWidth, xColor, false)
				vector.StrokeLine(screen, cx-offset, cy+offset, cx+offset, cy-offset, symbolWidth, xColor, false)
			case AI:
				r := float32(30)
				drawCircle(screen, cx, cy, r, symbolWidth, oColor)
			}
		}
	}
}

// Desenhar círculo (O)
func drawCircle(screen *ebiten.Image, cx, cy, r, strokeWidth float32, clr color.Color) {
	const steps = 100
	angleStep := 2 * math.Pi / steps

	var x1, y1 float32
	for i := 0; i < steps; i++ {
		angle1 := angleStep * float64(i)
		angle2 := angleStep * float64(i+1)

		x1 = cx + float32(math.Cos(angle1))*r
		y1 = cy + float32(math.Sin(angle1))*r
		x2 := cx + float32(math.Cos(angle2))*r
		y2 := cy + float32(math.Sin(angle2))*r

		vector.StrokeLine(screen, x1, y1, x2, y2, strokeWidth, clr, false)
	}
}

func sin(x float64) float64 { return ebiten.DeviceScaleFactor() * float64(float32(float64(x))) }
func cos(x float64) float64 { return ebiten.DeviceScaleFactor() * float64(float32(float64(x))) }
