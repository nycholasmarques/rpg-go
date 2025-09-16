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

type EntityType int

const (
	EntityTypeEnemy EntityType = iota
	EntityTypeHouse
	EntityTypeNPC
	EntityTypeTreasure
)

const speed = 1.4
const mapWidth = 50 * tileSize
const mapHeight = 50 * tileSize

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
	Type        EntityType
	Color       color.RGBA
	Sprite      *ebiten.Image
	OnCollision func(*Game)
}

func (e *Entity) Rect() (float64, float64, float64, float64) {
	return e.X, e.Y, e.W, e.H
}

func NewEbitenGameExploration(gs *model.GameState, screen GameActualScreen) *Game {
	LoadTiles()
	InitMap()

	playerImg := ebiten.NewImage(16, 16)
	playerImg.Fill(color.RGBA{255, 0, 0, 255})

	objects := []Entity{
		{
			X: 200, Y: 200, W: 16, H: 16,
			Type: EntityTypeEnemy,
			Color: color.RGBA{0, 0, 255, 255},
			OnCollision: func(g *Game) {
				g.Screen = ScreenBattle
				fmt.Println("começou batalha com inimigo")
			},
		},
		{
			X: 240, Y: 180, W: 16, H: 16,
			Type: EntityTypeEnemy,
			Color: color.RGBA{0, 0, 255, 255},
			OnCollision: func(g *Game) {
				g.Screen = ScreenBattle
				fmt.Println("começou batalha com inimigo")
			},
		},
		{
			X: 80, Y: 80, W: 16, H: 16,
			Type: EntityTypeHouse,
			Color: color.RGBA{0, 0, 200, 200},
			OnCollision: func(g *Game) {
				fmt.Println("entrou na casa")
			},
		},
		{
			X: 120, Y: 80, W: 16, H: 16,
			Type: EntityTypeHouse,
			Color: color.RGBA{0, 0, 200, 200},
			OnCollision: func(g *Game) {
				fmt.Println("entrou na casa")
			},
		},
		{
			X: 100, Y: 100, W: 16, H: 16,
			Type: EntityTypeNPC,
			Color: color.RGBA{0, 0, 150, 150},
			OnCollision: func(g *Game) {
				fmt.Println("dialogo com npc: Bem-vindo ao vilarejo!")
			},
		},
		{
			X: 90, Y: 120, W: 16, H: 16,
			Type: EntityTypeNPC,
			Color: color.RGBA{0, 0, 150, 150},
			OnCollision: func(g *Game) {
				fmt.Println("dialogo com npc: Cuidado na floresta!")
			},
		},
		{
			X: 400, Y: 400, W: 16, H: 16,
			Type: EntityTypeTreasure,
			Color: color.RGBA{255, 215, 0, 255},
			OnCollision: func(g *Game) {
				fmt.Println("encontrou um tesouro!")
			},
		},
	}

	for i := range objects {
		obj := &objects[i]
		obj.Sprite = ebiten.NewImage(int(obj.W), int(obj.H))
		obj.Sprite.Fill(obj.Color)
	}

	return &Game{
		GameState: gs,
		Screen:    screen,
		PlayerImg: playerImg,
		Player: Entity{
			X: gs.PosX,
			Y: gs.PosY,
			W: 16,
			H: 16,
		},
		Objects: objects,
	}
}

func (g *Game) Update() error {
	switch g.Screen {
	case ScreenMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Screen = ScreenExploration
		}
	case ScreenExploration:
		dx, dy := 0.0, 0.0
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			dx = -speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			dx = speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			dy = -speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
			dy = speed
		}

		if dx != 0 && dy != 0 {
			dx *= 0.707
			dy *= 0.707
		}

		newX := g.GameState.PosX + dx
		newY := g.GameState.PosY + dy
		tileX := int(newX / float64(tileSize))
		tileY := int(newY / float64(tileSize))
		if tileX >= 0 && tileX < len(worldMap[0]) && tileY >= 0 && tileY < len(worldMap) {
				if worldMap[tileY][tileX] != TILE_BLUE_AQUA && worldMap[tileY][tileX] != TILE_PINE_TREE {
						g.GameState.PosX = newX
						g.GameState.PosY = newY
				}
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
	switch g.Screen {
	case ScreenExploration:
		screenW, screenH := g.Layout(0, 0)
		camX := g.Player.X - float64(screenW)/2
		camY := g.Player.Y - float64(screenH)/2

		startX := int(camX/float64(tileSize)) - 1
		startY := int(camY/float64(tileSize)) - 1
		endX := startX + int(float64(screenW)/float64(tileSize)) + 2
		endY := startY + int(float64(screenH)/float64(tileSize)) + 2

		if startX < 0 {
			startX = 0
		}
		if startY < 0 {
			startY = 0
		}
		if endX > len(worldMap[0]) {
			endX = len(worldMap[0])
		}
		if endY > len(worldMap) {
			endY = len(worldMap)
		}

		for y := startY; y < endY; y++ {
			for x := startX; x < endX; x++ {
				tileIndex := worldMap[y][x]
				if tileIndex < 0 || tileIndex >= len(tiles) {
					continue
				}
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize)-camX, float64(y*tileSize)-camY)
				screen.DrawImage(tiles[tileIndex], op)
			}
		}

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(screenW)/2, float64(screenH)/2)
		screen.DrawImage(g.PlayerImg, opts)

		for _, obj := range g.Objects {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(obj.X-camX, obj.Y-camY)
			screen.DrawImage(obj.Sprite, o)
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf("explore | FPS: %.2f TPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}