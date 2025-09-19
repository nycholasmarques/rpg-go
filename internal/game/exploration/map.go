package exploration

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const mapPath = "assets/map/map-1.tmx"

func InitMap() (*ebiten.Image, *tiled.Map) {
    gameMap, err := tiled.LoadFile(mapPath)
    if err != nil {
        panic(err)
    }

    renderer, err := render.NewRenderer(gameMap)
    if err != nil {
        panic(err)
    }

    _ = renderer.RenderLayer(0)

    m := ebiten.NewImageFromImage(renderer.Result)
    return m, gameMap
}
