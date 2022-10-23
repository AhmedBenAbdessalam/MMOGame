package asset

import (
	"encoding/json"
	"errors"
	"image"
	_ "image/png"
	"io/fs"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/unitoftime/packer"
)

type Load struct {
	filesystem fs.FS
}

func NewLoad(filesystem fs.FS) *Load {
	return &Load{filesystem}
}

func (load *Load) Open(path string) (fs.File, error) {
	return load.filesystem.Open(path)
}

func (load *Load) Image(path string) (image.Image, error) {
	file, err := load.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (load *Load) Sprite(path string) (*pixel.Sprite, error) {

	img, err := load.Image(path)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)

	return pixel.NewSprite(pic, pic.Bounds()), nil

}

func (load *Load) Json(path string, dat interface{}) error {
	file, err := load.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, &dat)
}

func (load *Load) SpriteSheet(path string) (*SpriteSheet, error) {
	// Load the Json
	serializedSpritesheet := packer.SerializedSpritesheet{}
	err := load.Json(path, &serializedSpritesheet)
	if err != nil {
		return nil, err
	}

	// Load the image

	img, err := load.Image(serializedSpritesheet.ImageName)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	// Create the spritesheet
	bounds := pic.Bounds()
	lookup := make(map[string]*pixel.Sprite)
	for k, v := range serializedSpritesheet.Frames {
		rect := pixel.R(
			v.Frame.X,
			bounds.H()-v.Frame.Y,
			v.Frame.X+v.Frame.W,
			bounds.H()-(v.Frame.Y+v.Frame.H)).Norm()

		lookup[k] = pixel.NewSprite(pic, rect)
	}
	return NewSpritesheet(pic, lookup), nil
}

type SpriteSheet struct {
	picture pixel.Picture
	lookup  map[string]*pixel.Sprite
}

func NewSpritesheet(pic pixel.Picture, lookup map[string]*pixel.Sprite) *SpriteSheet {
	return &SpriteSheet{
		pic,
		lookup,
	}
}

func (s *SpriteSheet) Get(name string) (*pixel.Sprite, error) {
	sprite, ok := s.lookup[name]
	if !ok {
		return nil, errors.New("invalid sprite name")
	}
	return sprite, nil
}

func (s *SpriteSheet) Picture() pixel.Picture {
	return s.picture
}
