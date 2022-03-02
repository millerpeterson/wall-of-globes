package tiling

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/player"
)

// Tile - A rectangle offset in space.
type Tile struct {
	Offset geom.Vec  `json:"offset"`
	Rect   geom.Rect `json:"rect"`
}

// Tiling - A collection of named tiles.
type Tiling map[string]Tile

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
func Crops(srcVideo geom.Rect, wlt Tiling) map[string]player.Args {
	wlBound := EnclosingRect(wlt)
	fit := geom.MaxInnerFit(wlBound, srcVideo)
	wlOffset := geom.CenteringOffset(fit, srcVideo)
	wScale := float32(fit.Width) / float32(wlBound.Width)
	hScale := float32(fit.Height) / float32(wlBound.Height)
	crops := map[string]player.Args{}
	for tlName, tl := range wlt {
		l := wlOffset.X + int(float32(tl.Offset.X)*wScale)
		r := srcVideo.Width - (l + int(float32(tl.Rect.Width)*wScale))
		t := wlOffset.Y + int(float32(tl.Offset.Y)*hScale)
		b := srcVideo.Height - (t + int(float32(tl.Rect.Height)*hScale))
		crops[tlName] = player.Args{Top: t, Bottom: b, Left: l, Right: r}
	}
	return crops
}
