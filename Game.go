package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

type topScroll struct {
	shot            *ebiten.Image
	player          *ebiten.Image
	background      *ebiten.Image
	backgroundXView int
	xloc            int
	yloc            int
	score           int
	speed           int
	bulletXLoc      int
	bulletYLoc      int
	temp            bool
}

func (demo *topScroll) Update() error {
	backgroundWidth := demo.background.Bounds().Dx()
	maxX := backgroundWidth * 2
	demo.backgroundXView -= 4
	demo.backgroundXView %= maxX

	demo.speed = 1
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
			demo.yloc += demo.speed - 7
		} else {
			demo.yloc += demo.speed - 4
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
			demo.yloc += demo.speed + 3
		} else {
			demo.yloc += demo.speed + 2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		demo.temp = true
		if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
			demo.bulletXLoc += 4
		} else {
			demo.bulletXLoc = demo.xloc + 65
			demo.bulletYLoc = demo.yloc + 32
		}
	} else {
		if demo.bulletXLoc < 1000 {
			demo.bulletXLoc += 3
			//println(demo.bulletXLoc)
		}

	}
	//demo.bulletXLoc += 3
	//if demo.bulletXLoc > 1000 {
	//	demo.bulletXLoc = 0
	//
	//}

	return nil
}

func (demo *topScroll) Draw(screen *ebiten.Image) {
	drawOps := ebiten.DrawImageOptions{}
	const repeat = 3
	backgroundWidth := demo.background.Bounds().Dx()
	for count := 0; count < repeat; count += 1 {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(backgroundWidth*count), float64(-1000))
		drawOps.GeoM.Translate(float64(demo.backgroundXView), 0)
		screen.DrawImage(demo.background, &drawOps)
	}
	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(float64(demo.xloc), float64(demo.yloc))
	screen.DrawImage(demo.player, &drawOps)

	//for i := 0; i < 10; i += 1 {
	//	if demo.temp == true {
	//		println(i)
	//	}
	//}

	//for i := 0; i < 10; i += 1 {
	//	drawOps.GeoM.Reset()
	//	if demo.temp == true {
	//		drawOps.GeoM.Reset()
	//		drawOps.GeoM.Translate(float64(demo.bulletXLoc), float64(demo.bulletYLoc))
	//		screen.DrawImage(demo.shot, &drawOps)
	//	}

	//drawOps.GeoM.Reset()
	//drawOps.GeoM.Translate(float64(demo.bulletXLoc-i*10), float64(demo.bulletYLoc))
	//screen.DrawImage(demo.shot, &drawOps)
	//}

	if demo.temp == true {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(demo.bulletXLoc), float64(demo.bulletYLoc))
		screen.DrawImage(demo.shot, &drawOps)
	}
}

func (s topScroll) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 800)
	ebiten.SetWindowTitle("Game Project 1")

	playerPict, _, err := ebitenutil.NewImageFromFile("plane.png")
	if err != nil {
		fmt.Println("Unable to load image:", err)
	}
	backgroundPict, _, err := ebitenutil.NewImageFromFile("SpaceRed.png")
	if err != nil {
		fmt.Print("Can't load background:", err)
	}
	shotPict, _, err := ebitenutil.NewImageFromFile("M1.png")
	if err != nil {
		fmt.Print("Can't load shot:", err)
	}
	demo := topScroll{
		background: backgroundPict, player: playerPict,
		xloc: 1, yloc: 350, shot: shotPict,
		//bulletXLoc: 10, bulletYLoc: 350,
	}
	err = ebiten.RunGame(&demo)
	if err != nil {
		fmt.Print("Failed to run game:", err)
	}
}
