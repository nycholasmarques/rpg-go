package exploration

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nycholasmarques/rpg-go/internal/game/model"
)

type GameActualScreen int

const (
	ScreenMenu GameActualScreen = iota
	ScreenExploration
	ScreenBattle
)

const speed = 1.2

type Game struct {
	GameState *model.GameState
	Screen    GameActualScreen
	PlayerImg *ebiten.Image
	Objects   []Entity
	Player    Entity
}

type Entity struct {
	X, Y        float64
	W, H        float64
	OnCollision func(*Game)
}

func (e *Entity) Rect() (float64, float64, float64, float64) {
	return e.X, e.Y, e.W, e.H
}

func NewEbitenGameExploration(gs *model.GameState, screen GameActualScreen) *Game {
	// TODO: after, implement any pixelart based of class
	PlayerImg := ebiten.NewImage(16, 16)
	PlayerImg.Fill(color.RGBA{255, 0, 0, 255})
	return &Game{
		GameState: gs,
		Screen:    screen,
		PlayerImg: PlayerImg,
		Player: Entity{
			X: gs.PosX,
			Y: gs.PosY,
			W: 16,
			H: 16,
		},
		Objects: []Entity{
			{
				X: 50, Y: 50, W: 16, H: 16,
				OnCollision: func(g *Game) {
					fmt.Println("colidiu com objeto")
				},
			},
		},
	}
}

func (g *Game) Update() error {
	switch g.Screen {
	case ScreenMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Screen = ScreenExploration
		}
	case ScreenExploration:
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			g.GameState.PosX -= speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			g.GameState.PosX += speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			g.GameState.PosY -= speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			g.GameState.PosY += speed
		}
	}

	g.Player.X = g.GameState.PosX
	g.Player.Y = g.GameState.PosY

	for _, obj := range g.Objects {
		if checkCollision(g.Player, obj) {
			obj.OnCollision(g)
		}
	}
	return nil
}

func checkCollision(a, b Entity) bool {
	return a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: create an global menu
	switch g.Screen {
	case ScreenMenu:
		ebitenutil.DebugPrint(screen, "menu")
	case ScreenExploration:
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(g.GameState.PosX), float64(g.GameState.PosY))
		screen.DrawImage(g.PlayerImg, opts)

		for _, obj := range g.Objects {
			// TODO: change color to pixelart
			objImg := ebiten.NewImage(int(obj.W), int(obj.H))
			objImg.Fill(color.RGBA{0, 0, 255, 255})
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(obj.X, obj.Y)
			screen.DrawImage(objImg, o)
		}

		ebitenutil.DebugPrint(screen, "explore")
	case ScreenBattle:
		ebitenutil.DebugPrint(screen, "battle")
	}
	ebitenutil.DebugPrint(screen, g.GameState.DebugMsg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
