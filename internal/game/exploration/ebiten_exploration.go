package exploration

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
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

const tileSize = 16
const speed = 1.4

type Game struct {
	GameState *model.GameState
	Screen    GameActualScreen
	PlayerImg *ebiten.Image
	Objects   []Entity
	Player    Entity

	inDialog   bool
	DialogText []string
	Map        *ebiten.Image
	GameMap    *tiled.Map
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
	m, gm := InitMap()

	playerImg := ebiten.NewImage(16, 16)
	playerImg.Fill(color.RGBA{255, 0, 0, 255})

	objects := []Entity{
		{
			X: 200, Y: 200, W: 16, H: 16,
			Type:  EntityTypeEnemy,
			Color: color.RGBA{0, 0, 255, 255},
			OnCollision: func(g *Game) {
				g.Screen = ScreenBattle
				fmt.Println("começou batalha com inimigo")
			},
		},
		{
			X: 100, Y: 100, W: 16, H: 16,
			Type:  EntityTypeNPC,
			Color: color.RGBA{0, 0, 150, 150},
			OnCollision: func(g *Game) {
				g.inDialog = true
				g.DialogText = []string{
					"Olá my friend, vi que tem monstros javascript por ai...",
					"Eu vi eles na parte de baixo do mapa.",
					"... cuidado!",
				}
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
		Map:     m,
		GameMap: gm,
	}
}

var countDialogNPC = 0
var lastDialogAdvance time.Time

const dialogCooldown = 200 * time.Millisecond

func (g *Game) Update() error {
	switch g.Screen {
	case ScreenMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Screen = ScreenExploration
		}
	case ScreenExploration:
		if g.inDialog {
			if (ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeySpace)) &&
				time.Since(lastDialogAdvance) > dialogCooldown {

				lastDialogAdvance = time.Now()
				countDialog := len(g.DialogText)

				if countDialogNPC < countDialog-1 {
					countDialogNPC++
				} else {
					g.inDialog = false
					g.DialogText = []string{}
					countDialogNPC = 0
					g.GameState.PosX = g.GameState.PosX - 2
					g.GameState.PosY = g.GameState.PosY - 2
				}
			}
			return nil
		}

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

		playerW, playerH := g.Player.W, g.Player.H

		if !g.IsSollidAt(newX, newY) &&
			!g.IsSollidAt(newX+playerW-1, newY) &&
			!g.IsSollidAt(newX, newY+playerH-1) &&
			!g.IsSollidAt(newX+playerW-1, newY+playerH-1) {

			g.GameState.PosX = newX
			g.GameState.PosY = newY
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

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-camX, -camY)
		screen.DrawImage(g.Map, op)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(screenW)/2, float64(screenH)/2)
		screen.DrawImage(g.PlayerImg, opts)

		for _, obj := range g.Objects {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Translate(obj.X-camX, obj.Y-camY)
			screen.DrawImage(obj.Sprite, o)
		}

		if g.inDialog {
			dialogBoxHeight := 80
			dialogBox := ebiten.NewImage(screenW, dialogBoxHeight)
			dialogBox.Fill(color.RGBA{0, 0, 0, 180})

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(0, float64(screenH-dialogBoxHeight))
			screen.DrawImage(dialogBox, opts)

			ebitenutil.DrawRect(screen, 0, float64(screenH-dialogBoxHeight), float64(screenW), 2, color.White)
			ebitenutil.DrawRect(screen, 0, float64(screenH-2), float64(screenW), 2, color.White)

			text := g.DialogText[countDialogNPC]
			maxWidth := screenW - 20
			lines := wrapText(text, maxWidth/8)
			for i, line := range lines {
				ebitenutil.DebugPrintAt(screen, line, 10, screenH-dialogBoxHeight+10+i*15)
			}

			if time.Now().Unix()%2 == 0 {
				ebitenutil.DebugPrintAt(screen, "▶", screenW-20, screenH-20)
			}
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf("explore | FPS: %.2f TPS: %.2f",
			ebiten.ActualFPS(), ebiten.ActualTPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) IsSollidAt(x, y float64) bool {
	tileX := int(x / float64(tileSize))
	tileY := int(y / float64(tileSize))

	if tileX < 0 || tileY < 0 || tileX >= g.GameMap.Width || tileY >= g.GameMap.Height {
		return true
	}

	layer := g.GameMap.Layers[0]
	tile := layer.Tiles[tileY*g.GameMap.Width+tileX]
	if tile == nil {
		return false
	}

	tsTile, _ := tile.Tileset.GetTilesetTile(tile.ID)
	if tsTile == nil {
		return false
	}

	for _, og := range tsTile.ObjectGroups {
		for _, obj := range og.Objects {
			rectX := float64(tileX*tileSize) + obj.X
			rectY := float64(tileY*tileSize) + obj.Y
			rectW := obj.Width
			rectH := obj.Height

			if x >= rectX && x < rectX+rectW &&
				y >= rectY && y < rectY+rectH {
				return true
			}
		}
	}
	
	if tsTile.Properties.GetBool("collision") {
		return true
	}

	return false
}

func wrapText(text string, maxChars int) []string {
	runes := []rune(text)
	var lines []string
	var line []rune

	for _, r := range runes {
		line = append(line, r)
		if len(line) >= maxChars && r == ' ' {
			lines = append(lines, string(line))
			line = []rune{}
		}
	}
	if len(line) > 0 {
		lines = append(lines, string(line))
	}
	return lines
}
