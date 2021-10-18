package main

import (
	"fmt"
	"image"
	_ "image/png"

	"./map"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

const (
	tileSize = 16
	tileXNum = 25
)

var (
	tilesImage *ebiten.Image
	sWidth int
	sHeight int
)

var terrain _map.Map

func init() {
	terrain, _ = _map.Read("./rsrc/scenarios/zakhazi/towns/town0.map")

	sWidth  = terrain.Length() * tileSize
	sHeight = terrain.Height() * tileSize

	file, _ := ebitenutil.OpenFile("./rsrc/graphics/termap.png")
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func (g *Game) Update (screen *ebiten.Image) error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return sWidth, sHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	xNum := sWidth / tileSize
	for _, r := range terrain.Map {
		{
			for i, t := range r {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

				sx := (t.Terrain % tileXNum) * tileSize
				sy := (t.Terrain / tileXNum) * tileSize
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
			}
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(sWidth*2, sHeight*2)
	ebiten.SetWindowTitle("Go Classic Blades of Exile")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}