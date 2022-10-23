package main

//go:generate  $GOPATH\pkg\mod\github.com\unitoftime\packer@v0.0.0-20220105185326-f541e031de11\cmd\packer\packer.exe --input images --stats

import (
	"os"

	"github.com/AhmedBenAbdessalam/MMOGame/engine/asset"
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

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))

		// input handling
		if win.Pressed(pixelgl.KeyLeft) {
			knightPosition.X -= 2.0
		}
		if win.Pressed(pixelgl.KeyRight) {
			knightPosition.X += 2.0
		}
		if win.Pressed(pixelgl.KeyUp) {
			knightPosition.Y += 2.0
		}
		if win.Pressed(pixelgl.KeyDown) {
			knightPosition.Y -= 2.0
		}

		knightSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(knightPosition))
		win.Update()
	}
}
