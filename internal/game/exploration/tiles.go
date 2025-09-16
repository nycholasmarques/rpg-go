package exploration

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nycholasmarques/rpg-go/assets"
)

var (
	tiles    []*ebiten.Image
	worldMap [][]int
)

const tileSize = 16

const (
	TILE_BLUE = iota
	TILE_BRICK_GREY
	TILE_BRICK_GREY2
	TILE_BRICK_GREY3
	TILE_BRICK_GREY4
	TILE_BRICK_RED1
	TILE_BRICK_RED2
	TILE_BRICK_RED3
	TILE_BRICK_RED4
	TILE_BRICK_WITH_BROWN_SAND
	TILE_RED
	TILE_SAND
	TILE_GRASS
	TILE_GRASS_WITH_FLOWERS
	TILE_BLUE_AQUA
	TILE_BRICK
	TILE_ROCK_BROWN
	TILE_WOOD_FLOOR
	TILE_RED_FLOOR
	TILE_SAND_CIRCLE_FORMAT
	TILE_GRASS_CIRCLE_FORMAT
	TILE_BUSH
	TILE_ICE
	TILE_BLACK
	TILE_ROCK_BROWN_WITH_SHADOW
	TILE_LEFT_AND_TOP_WALL_BROWN
	TILE_TOP_WALL_BROWN
	TILE_RIGHT_AND_TOP_WALL_BROWN
	TILE_CANISTER_BROWN
	TILE_CANISTER_BROWN_BROKEN
	TILE_ICE_LINES1
	TILE_PINE_TREE
	TILE_WELL
	TILE_LEFT_WALL_BROWN
	TILE_DARK_PATH_DOOR
	TILE_RIGHT_WALL_BROWN
	TILE_TRUNK_OPEN
	TILE_TRUNK_CLOSED
	TILE_ICE_LINES2
	TILE_TREE
	TILE_PLATAFORM_ROUNDED
	TILE_LEFT_DOWN_WALL_BROWN
	TILE_DOWN_WALL_BROWN
	TILE_DOWN_RIGHT_WALL_BROWN
	TILE_MUSHROOMS
	TILE_MUSHROOM2
	TILE_BED_PILLOW_PART
	TILE_ARMCHAIR
	TILE_EAGLE_SCULPTURE
	TILE_DOOR_RAINBOW
	TILE_DOOR_JAIL
	TILE_DOOR_EMPTY_BLACK
	TILE_TWO_CONES_HORIZONTAL
	TILE_BED_RAINBOW_FLOOR
	TILE_BED_REST_PART
	TILE_TABLE
	TILE_DRAWERS
	TILE_SIDE_LADDER
	TILE_STAIRS_DOWN
	TILE_ROCK_GREY_FLOOR
	TILE_TWO_CONES_RAINBOW_VERTICAL
	TILE_TORCH1
	TILE_TORCH2
	TILE_SAND_HILL
	TILE_SAND_GREY_HILL
	TILE_GRASS_VARIATION1
	TILE_GRAS_VARIATION2
	TILE_ROCK_BROWN_ON_FIRE
	TILE_PLATE
	TILE_TORCH_BROKEN
	TILE_PRECIOUS_ROCKS
	TILE_PLATE_WARRIOR
	TILE_PLATE_MOON_AND_STARS
	TILE_MUD_TERRAIN
	TILE_BRICK_FLOOW
	TILE_FIRE_LAVA1
	TILE_FIRE_LAVA2
	TILE_TREE_2
	TILE_TWO_PINE_TREE
	TILE_PLATE_SWORD
	TILE_PLATE_PORTION
	TILE_NARROW_VERTICAL_PATH_GREY
	TILE_NARROW_HORIZONTAL_PATH_GREY
	EMPTY1
	EMPTY2
	EMPTY3
	EMPTY4
	EMPTY5
	EMPTY6
	TILE_NARROW_VERTICAL_PATH_GOLD
	TILE_NARROW_HORIZONTAL_PATH_GOLD
	EMPTY7
	TILE_PILLAR_TOP_PART
	TILE_SOME_TREE_TOP_VISION
	EMPTY8
	EMPTY9
	EMPTY10
	EMPTY11
	EMPTY12
	EMPTY13
	TILE_PILLAR_DOWN_PART
)

func LoadTiles() {
	tiles = []*ebiten.Image{}

	f2, err := assets.Water_Middle.Open("exploration/Water_Middle.png")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	img2, _, err := image.Decode(f2)
	if err != nil {
		panic(err)
	}
	tiles = append(tiles, ebiten.NewImageFromImage(img2))

	f3, err := assets.Basictiles.Open("exploration/basictiles.png")
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	img3, _, err := image.Decode(f3)
	if err != nil {
		panic(err)
	}
	tileset := ebiten.NewImageFromImage(img3)

	bounds := tileset.Bounds()
	for y := 0; y < bounds.Dy(); y += tileSize {
		for x := 0; x < bounds.Dx(); x += tileSize {
			rect := image.Rect(x, y, x+tileSize, y+tileSize)
			sub := tileset.SubImage(rect).(*ebiten.Image)
			tiles = append(tiles, sub)
		}
	}
}

func InitMap() {
	worldMap = make([][]int, 50)
	for i := range worldMap {
		worldMap[i] = make([]int, 50)
		for j := range worldMap[i] {
			worldMap[i][j] = TILE_GRASS_VARIATION1
		}
	}

	for y := 2; y < 12; y++ {
		for x := 2; x < 12; x++ {
			if x == 2 {
				worldMap[x-1][y] = TILE_BRICK_GREY
			}
			worldMap[y][x] = TILE_BRICK
		}
	}
	worldMap[7][7] = TILE_WELL

	for x := 12; x < 30; x++ {
		worldMap[7][x] = TILE_SAND
	}
	for y := 7; y < 40; y++ {
		worldMap[y][29] = TILE_SAND
	}

	for y := 35; y < 45; y++ {
		for x := 35; x < 45; x++ {
			worldMap[y][x] = TILE_BLUE_AQUA
		}
	}
	for y := 34; y < 46; y++ {
		worldMap[y][34] = TILE_SAND
		worldMap[y][45] = TILE_SAND
	}
	for x := 34; x < 46; x++ {
		worldMap[34][x] = TILE_SAND
		worldMap[45][x] = TILE_SAND
	}

	worldMap[15][15] = TILE_TRUNK_CLOSED
	worldMap[20][20] = TILE_ROCK_BROWN
	worldMap[25][25] = TILE_MUSHROOMS
}