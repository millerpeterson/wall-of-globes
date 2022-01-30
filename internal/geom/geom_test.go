package geom

import "testing"

func TestMaxInnerFit(t *testing.T) {
	// Same aspect ratio
	inner := Rect{300, 200}
	outer := Rect{600, 400}
	expected := outer
	max := MaxInnerFit(inner, outer)
	if max != expected {
		t.Errorf("Expect %v, got %v", expected, max)
	}

	// Inner wider ratio than outer
	inner = Rect{400, 200}
	outer = Rect{600, 400}
	expected = Rect{600, 300}
	max = MaxInnerFit(inner, outer)
	if max != expected {
		t.Errorf("Expect %v, got %v", expected, max)
	}

	// Inner narrower ratio than outer
	inner = Rect{200, 200}
	outer = Rect{600, 400}
	expected = Rect{400, 400}
	max = MaxInnerFit(inner, outer)
	if max != expected {
		t.Errorf("Expect %v, got %v", expected, max)
	}
}
