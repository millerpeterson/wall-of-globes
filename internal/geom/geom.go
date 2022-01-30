package geom

type Rect struct {
	Width  int
	Height int
}

func (r Rect) AspectRatio() float32 {
	return float32(r.Width) / float32(r.Height)
}

// MaxInnerFit Return the largest (i.e. max area) rectangle with
// the same aspect ratio of `inner` that fits inside `outer`.
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
