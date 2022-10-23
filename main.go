package main

//go:generate  $GOPATH\pkg\mod\github.com\unitoftime\packer@v0.0.0-20220105185326-f541e031de11\cmd\packer\packer.exe --input images --stats

import (
	"os"

	"github.com/AhmedBenAbdessalam/MMOGame/engine/asset"
	"github.com/AhmedBenAbdessalam/MMOGame/engine/render"
	"github.com/AhmedBenAbdessalam/MMOGame/engine/tilemap"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {

	cfg := pixelgl.WindowConfig{
		Title:     "MMO",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)
	load := asset.NewLoad(os.DirFS("./"))

	spritesheet, err := load.SpriteSheet("packed.json")
	if err != nil {
		panic(err)
	}

	knightSprite, err := spritesheet.Get("knight.png")
	if err != nil {
		panic(err)
	}

	knightPosition := win.Bounds().Center()

	knight2Sprite, err := spritesheet.Get("knight2.png")
	if err != nil {
		panic(err)
	}

	knight2Position := win.Bounds().Center()

	people := make([]Person, 0)

	people = append(people, NewPerson(knightSprite, knightPosition, Keybinds{
		Up:    pixelgl.KeyUp,
		Down:  pixelgl.KeyDown,
		Left:  pixelgl.KeyLeft,
		Right: pixelgl.KeyRight,
	}))
	people = append(people, NewPerson(knight2Sprite, knight2Position, Keybinds{
		Up:    pixelgl.KeyW,
		Down:  pixelgl.KeyS,
		Left:  pixelgl.KeyA,
		Right: pixelgl.KeyD,
	}))

	grassSprite, err := spritesheet.Get("grass.png")
	if err != nil {
		panic(err)
	}
	tileSize := 16
	mapSize := 100
	tiles := make([][]tilemap.Tile, mapSize)
	for x := range tiles {
		tiles[x] = make([]tilemap.Tile, mapSize)
		for y := range tiles[x] {

			tiles[x][y] = tilemap.Tile{
				Type:   0,
				Sprite: grassSprite,
			}
		}
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet.Picture())
	tmap := tilemap.New(tiles, batch, tileSize)
	tmap.Rebatch()

	camera := render.NewCamera(win, 0, 0)
	zoomSpeed := 0.1

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		scroll := win.MouseScroll()
		if scroll.Y != 0 {
			camera.Zoom += zoomSpeed * scroll.Y
		}

		for i := range people {
			people[i].HandleInput(win)
		}

		camera.Position = people[0].Position
		camera.Update()

		win.SetMatrix(camera.Mat())
		tmap.Draw(win)
		for i := range people {
			people[i].Draw(win)

		}
		win.SetMatrix(pixel.IM)

		win.Update()
	}
}

type Keybinds struct {
	Up, Down, Left, Right pixelgl.Button
}

type Person struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
	Keybinds Keybinds
}

func NewPerson(sprite *pixel.Sprite, position pixel.Vec, keybinds Keybinds) Person {
	return Person{
		sprite,
		position,
		keybinds,
	}
}

func (p *Person) Draw(win *pixelgl.Window) {
	p.Sprite.Draw(win, pixel.IM.Moved(p.Position))
}

func (p *Person) HandleInput(win *pixelgl.Window) {

	if win.Pressed(p.Keybinds.Up) {
		p.Position.Y += 2.0
	}
	if win.Pressed(p.Keybinds.Down) {
		p.Position.Y -= 2.0
	}
	if win.Pressed(p.Keybinds.Left) {
		p.Position.X -= 2.0
	}
	if win.Pressed(p.Keybinds.Right) {
		p.Position.X += 2.0
	}

}
