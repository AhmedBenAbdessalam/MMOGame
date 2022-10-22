package main

import (
	"image"
	_ "image/png"
	"os"

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

	knightSprite, err := getSprite("knight.png")
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

func getSprite(path string) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)

	return pixel.NewSprite(pic, pic.Bounds()), nil

}
