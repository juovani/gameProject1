package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"math/rand"
)

type topScroll struct {
	player          *ebiten.Image
	background      *ebiten.Image
	bullet          *ebiten.Image
	enemy           *ebiten.Image
	backgroundXView int
	xloc            int
	yloc            int
	score           int
	speed           int
	temp            bool
	evil            []enemys
	shot            []shots
}

type shots struct {
	shot       *ebiten.Image
	bulletXLoc int
	bulletYLoc int
}
type enemys struct {
	enemy     *ebiten.Image
	enemyXLoc int
	enemyYLoc int
}

func newShots(image *ebiten.Image) shots {
	return shots{
		shot: image,
	}
}

func newEnemy(MAxHeight int, image *ebiten.Image) enemys {
	return enemys{
		enemy:     image,
		enemyXLoc: 1000,
		enemyYLoc: rand.Intn(MAxHeight),
	}

}

func (demo *topScroll) Update() error {
	backgroundWidth := demo.background.Bounds().Dx()
	maxX := backgroundWidth * 2
	demo.backgroundXView -= 4
	demo.backgroundXView %= maxX

	demo.speed = 1
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		demo.yloc += demo.speed - 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		demo.yloc += demo.speed + 2
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		demo.temp = true

		newBullet := newShots(demo.bullet)
		newBullet.bulletXLoc = demo.xloc + 65
		newBullet.bulletYLoc = demo.yloc + 32
		demo.shot = append(demo.shot, newBullet)
	} else {
		for i := range demo.shot {
			demo.shot[i].bulletXLoc += 4
		}
	}

	randomNumber := rand.Intn(500)

	if randomNumber < 5 {
		newEnemy := newEnemy(900, demo.enemy)
		demo.evil = append(demo.evil, newEnemy)
	}

	for i := range demo.evil {
		demo.evil[i].enemyXLoc -= 2
	}
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
	for _, enemy := range demo.evil {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(enemy.enemyXLoc), float64(enemy.enemyYLoc))
		screen.DrawImage(enemy.enemy, &drawOps)
	}
	if demo.temp == true {
		for _, shot := range demo.shot {
			drawOps.GeoM.Reset()
			drawOps.GeoM.Translate(float64(shot.bulletXLoc), float64(shot.bulletYLoc))
			screen.DrawImage(shot.shot, &drawOps)
		}
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
	enemyPict, _, err := ebitenutil.NewImageFromFile("enemy.png")
	if err != nil {
		fmt.Print("Can't load enemy:", err)
	}

	allShots := make([]shots, 0, 20)
	allEnemys := make([]enemys, 0, 15)
	for i := 0; i < 5; i += 1 {
		allEnemys = append(allEnemys, newEnemy(900, enemyPict))
	}
	demo := topScroll{
		background: backgroundPict, player: playerPict,
		xloc: 1, yloc: 350,
		shot:   allShots,
		evil:   allEnemys,
		bullet: shotPict,
		enemy:  enemyPict,
	}
	err = ebiten.RunGame(&demo)
	if err != nil {
		fmt.Print("Failed to run game:", err)
	}
}
