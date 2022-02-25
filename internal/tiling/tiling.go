package tiling

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/player"
)

// Tile - A rectangle offset in space.
type Tile struct {
	Offset geom.Vec
	Rect   geom.Rect
}

// Tiling - A collection of tiles.
type Tiling []Tile

// EnclosingRect - the minimal rectangle that covers the tiles in a wall.
func EnclosingRect(wlt Tiling) geom.Rect {
	var maxRight int
	var maxBottom int
	for _, tl := range wlt {
		right := tl.Offset.X + tl.Rect.Width
		if right > maxRight {
			maxRight = right
		}
		bottom := tl.Offset.Y + tl.Rect.Height
		if bottom > maxBottom {
			maxBottom = bottom
		}
	}
	return geom.Rect{Width: maxRight, Height: maxBottom}
}

// Crops - get an array of crops for `srcVideo`, each one corresponding
// to a tile in a wall.
func Crops(srcVideo geom.Rect, wlt Tiling) (crops []player.Args) {
	wlBound := EnclosingRect(wlt)
	fit := geom.MaxInnerFit(wlBound, srcVideo)
	wlOffset := geom.CenteringOffset(fit, srcVideo)
	for _, tl := range wlt {
		l := wlOffset.X + tl.Offset.X
		r := srcVideo.Width - (l + tl.Rect.Width)
		t := wlOffset.Y + tl.Offset.Y
		b := srcVideo.Height - (t + tl.Rect.Height)
		crops = append(crops, player.Args{Top: t, Bottom: b, Left: l, Right: r})
	}
	return crops
}
