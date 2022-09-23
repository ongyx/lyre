package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento/ecs"
	"github.com/ongyx/lyre"
)

type Game struct {
	*ecs.World
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return 256, 256
}

func main() {
	if err := ebiten.RunGame(&Game{lyre.Scene()}); err != nil {
		panic(err)
	}
}
