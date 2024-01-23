// Copyright 2021 zs. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowPosition(ebiten.ScreenSizeInFullscreen())
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetTPS(30)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
