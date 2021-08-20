// Copyright 2020 zs. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed dot.png
var bg []byte

var dots *ebiten.Image
var offImg *ebiten.Image
var onImg *ebiten.Image
var screenWidth int
var screenHeight int

const (
	dotWidth = 36
	offset   = 0
)

func init() {

	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()

	img, err := png.Decode(bytes.NewReader(bg))
	if err != nil {
		log.Fatal(err)
	}
	dots = ebiten.NewImageFromImage(img)
	offImg = dots.SubImage(image.Rect(0, 0, dotWidth, dotWidth)).(*ebiten.Image)
	onImg = dots.SubImage(image.Rect(dotWidth, 0, dotWidth*2, dotWidth)).(*ebiten.Image)
}

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	timestr := time.Now().Format("15:04:05")
	x := (screenWidth - (dotWidth*8*6 - dotWidth)) / 2
	y := (screenHeight - dotWidth*7) / 2
	for i := 0; i < 8; i++ {
		char := timestr[i : i+1]
		for j := 0; j < 7; j++ {
			for k := 0; k < 5; k++ {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x+k*dotWidth-j*offset), float64(y+j*dotWidth))
				if font[char][j][k] == 0 {
					screen.DrawImage(offImg, op)
				} else {
					screen.DrawImage(onImg, op)
				}
			}
		}
		x += dotWidth * 6
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
