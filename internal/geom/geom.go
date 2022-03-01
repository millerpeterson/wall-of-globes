package geom

type Vec struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Rect struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (r Rect) AspectRatio() float32 {
	return float32(r.Width) / float32(r.Height)
}

// MaxInnerFit - Return the largest (i.e. max area) rectangle with the same
// aspect ratio of `inner` that fits inside `outer`.
func MaxInnerFit(inner Rect, outer Rect) Rect {
	innerAspect := inner.AspectRatio()
	// Try fitting to width
	widthFitHeight := int(float32(outer.Width) * (1 / innerAspect))
	if widthFitHeight <= outer.Height {
		return Rect{outer.Width, widthFitHeight}
	}
	// Fit to height
	heightFitWidth := int(float32(outer.Height) * innerAspect)
	return Rect{heightFitWidth, outer.Height}
}

// MapRectPos - Return the position that `srcPos` from `src` would map to in
// `target`, assuming a linear mapping.
func MapRectPos(src Rect, target Rect, srcPos Vec) Vec {
	widthRatio := float32(target.Width) / float32(src.Width)
	heightRatio := float32(target.Height) / float32(src.Height)
	return Vec{
		X: int(float32(srcPos.X) * widthRatio),
		Y: int(float32(srcPos.Y) * heightRatio),
	}
}

// CenteringOffset - Return the offset that would put the center of inner in the center
// of outer.
func CenteringOffset(inner Rect, outer Rect) Vec {
	return Vec{
		X: (outer.Width - inner.Width) / 2,
		Y: (outer.Height - inner.Height) / 2,
	}
}
