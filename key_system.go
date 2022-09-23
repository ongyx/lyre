package lyre

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ongyx/bento"
	"github.com/ongyx/bento/ecs"

	"github.com/ongyx/lyre/assets"
)

const tilesize = 16

var (
	spritesheet = assets.NewResource("sprites/lyre.png")
	tilemap     = [][][]int{
		// notes bg
		{
			{-1, 3, 3, 3, 3, 3, 3, 3},
			{-1, 3, 3, 3, 3, 3, 3, 3},
			{-1, 3, 3, 3, 3, 3, 3, 3},
		},
		// notes
		{
			{-1, 8, 9, 10, 11, 12, 13, 14},
			{-1, 8, 9, 10, 11, 12, 13, 14},
			{-1, 8, 9, 10, 11, 12, 13, 14},
		},
		// clefs
		{
			{0, -1, -1, -1, -1, -1, -1, -1},
			{1, -1, -1, -1, -1, -1, -1, -1},
			{2, -1, -1, -1, -1, -1, -1, -1},
		},
	}

	bgColor = color.NRGBA{203, 219, 252, 255}
)

type KeySystem struct {
	view *ecs.View
	img  *ebiten.Image
}

func (k *KeySystem) Init(w *ecs.World) {
	k.view = ecs.NewView(w, ecs.Type[Key]())

	img, err := spritesheet.OpenImage()
	if err != nil {
		panic(err)
	}

	ts := bento.NewTileset(ebiten.NewImageFromImage(img), tilesize)

	// pre-render tilemap
	for _, m := range tilemap {
		i := ts.Render(m)

		if k.img != nil {
			k.img.DrawImage(i, nil)
		} else {
			k.img = i
		}
	}
}

func (k *KeySystem) Update(w *ecs.World) error {
	key := ecs.Query[Key](w)

	k.view.Each(func(e ecs.Entity) {
		k := key.Get(e)
		if inpututil.IsKeyJustPressed(k.key) {
			fmt.Printf("key pressed: %s\n", k.key)
		}
	})

	return nil
}

func (k *KeySystem) Render(w *ecs.World, i *ebiten.Image) {
	i.Fill(bgColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)

	i.DrawImage(k.img, op)
}
