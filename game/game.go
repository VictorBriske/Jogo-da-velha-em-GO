package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	Player       = 1
	AI           = -1
	Empty        = 0
	ScreenWidth  = 380
	ScreenHeight = 380
	CellSize     = 100
	BoardSize    = CellSize * 3
)

var (
	OffsetX = float32((ScreenWidth - BoardSize) / 2)
	OffsetY = float32((ScreenHeight - BoardSize) / 2)
)

type Game struct {
	board       [3][3]int
	turn        int
	gameOver    bool
	difficulty  int
	inMenu      bool
	inputLocked bool
	showIntro   bool
	introFrame  int
	ballX       float64
	textAlpha   float64
}

func NewGame() *Game {
	return &Game{
		board:       [3][3]int{},
		turn:        Player,
		inMenu:      false,
		inputLocked: false,
		showIntro:   true,
		introFrame:  0,
		ballX:       -50,
		textAlpha:   0.0,
	}
}

func (g *Game) Update() error {
	if g.showIntro {
		g.introFrame++
		g.ballX += 3
		if g.introFrame > 60 {
			g.textAlpha += 0.02
			if g.textAlpha > 1 {
				g.textAlpha = 1
			}
		}
		if g.introFrame > 180 {
			g.showIntro = false
			g.inMenu = true
		}
		return nil
	}

	if g.inputLocked {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.inputLocked = false
		}
		return nil
	}

	if g.inMenu {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 140 && x <= 240 && y >= 100 && y <= 140:
				g.difficulty = 1
				g.inMenu = false
				g.inputLocked = true
			case x >= 140 && x <= 240 && y >= 160 && y <= 200:
				g.difficulty = 2
				g.inMenu = false
				g.inputLocked = true
			case x >= 140 && x <= 240 && y >= 220 && y <= 260:
				g.difficulty = 3
				g.inMenu = false
				g.inputLocked = true
			}
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}

	if g.gameOver {
		x, y := ebiten.CursorPosition()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
			x >= 140 && x <= 240 && y >= 320 && y <= 360 {
			g.Reset()
		}
		return nil
	}

	if g.turn == Player {
		g.handlePlayerInput()
	} else {
		row, col := GetBestMove(g.board, g.difficulty)
		if g.board[row][col] == Empty {
			g.board[row][col] = AI
			g.turn = Player
		}
	}

	if winner := CheckWinner(g.board); winner != 0 || IsBoardFull(g.board) {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.showIntro {
		g.drawIntro(screen)
		return
	}

	if g.inMenu {
		drawButton(screen, 140, 100, "Fácil", color.RGBA{100, 200, 100, 255})
		drawButton(screen, 140, 160, "Médio", color.RGBA{200, 200, 100, 255})
		drawButton(screen, 140, 220, "Difícil", color.RGBA{200, 100, 100, 255})
		return
	}

	DrawBoard(screen, g.board)

	if g.gameOver {
		// Centraliza o botão "Reiniciar"
		btnLabel := "Reiniciar"
		btnWidth := 100
		btnX := (ScreenWidth - btnWidth) / 2
		drawButton(screen, btnX, 320, btnLabel, color.RGBA{80, 80, 80, 255})

		// Mensagem do resultado
		msg := "Empate!"
		if winner := CheckWinner(g.board); winner == Player {
			msg = "Você venceu!"
		} else if winner == AI {
			msg = "IA venceu!"
		}

		// Centraliza a mensagem
		msgX := (ScreenWidth - len(msg)*7) / 2
		text.Draw(screen, msg, basicfont.Face7x13, msgX, 310, color.White)
	}

}

func drawButton(screen *ebiten.Image, x, y int, label string, clr color.Color) {
	vector.DrawFilledRect(screen, float32(x), float32(y), 100, 40, clr, false)
	text.Draw(screen, label, basicfont.Face7x13, x+25, y+25, color.Black)
}

func (g *Game) drawIntro(screen *ebiten.Image) {
	screen.Fill(color.Black)
	vector.DrawFilledCircle(screen, float32(g.ballX), 190, 20, color.RGBA{100, 200, 255, 255}, false)

	alpha := uint8(g.textAlpha * 255)
	textStr := "Briske's TOE"
	textWidth := len(textStr) * 7 // 7 pixels por caractere na fonte básica
	x := (ScreenWidth - textWidth) / 2

	text.Draw(screen, textStr, basicfont.Face7x13, x, 250, color.RGBA{255, 255, 255, alpha})
}

func (g *Game) Reset() {
	g.board = [3][3]int{}
	g.turn = Player
	g.gameOver = false
	g.inMenu = true
	g.inputLocked = false
}

func (g *Game) handlePlayerInput() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		gridX := x - int(OffsetX)
		gridY := y - int(OffsetY)
		if gridX >= 0 && gridY >= 0 {
			col := gridX / CellSize
			row := gridY / CellSize
			if row < 3 && col < 3 && g.board[row][col] == Empty {
				g.board[row][col] = Player
				g.turn = AI
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
