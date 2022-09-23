package lyre

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/bento/ecs"
)

var keys = [...]ebiten.Key{
	ebiten.KeyQ,
	ebiten.KeyW,
	ebiten.KeyE,
	ebiten.KeyR,
	ebiten.KeyT,
	ebiten.KeyY,
	ebiten.KeyU,
	ebiten.KeyA,
	ebiten.KeyS,
	ebiten.KeyD,
	ebiten.KeyF,
	ebiten.KeyG,
	ebiten.KeyH,
	ebiten.KeyJ,
	ebiten.KeyZ,
	ebiten.KeyX,
	ebiten.KeyC,
	ebiten.KeyV,
	ebiten.KeyB,
	ebiten.KeyN,
}

func Scene() *ecs.World {
	w := ecs.NewWorld(32)

	key := ecs.Register[Key](w, len(keys))

	for _, k := range keys {
		e := w.Spawn()
		key.Insert(e, Key{k})
	}

	w.Register(&KeySystem{})

	return w
}
