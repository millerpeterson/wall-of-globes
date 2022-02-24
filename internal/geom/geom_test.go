package geom

import "testing"

func TestMaxInnerFit(t *testing.T) {
	t.Run("Same aspect ratio", func(t *testing.T) {
		inner := Rect{300, 200}
		outer := Rect{600, 400}
		expected := outer
		max := MaxInnerFit(inner, outer)
		if max != expected {
			t.Errorf("Expected %v, got %v", expected, max)
		}
	})

	t.Run("Inner wider ratio than outer", func(t *testing.T) {
		inner := Rect{400, 200}
		outer := Rect{600, 400}
		expected := Rect{600, 300}
		max := MaxInnerFit(inner, outer)
		if max != expected {
			t.Errorf("Expected %v, got %v", expected, max)
		}
	})

	t.Run("Inner narrower ratio than outer", func(t *testing.T) {
		inner := Rect{200, 200}
		outer := Rect{600, 400}
		expected := Rect{400, 400}
		max := MaxInnerFit(inner, outer)
		if max != expected {
			t.Errorf("Expected %v, got %v", expected, max)
		}
	})
}

func TestMapRectPos(t *testing.T) {
	t.Run("src smaller than target", func(t *testing.T) {
		src := Rect{300, 200}
		target := Rect{600, 400}
		srcPos := Vec{150, 100}
		expected := Vec{300, 200}
		targetPos := MapRectPos(src, target, srcPos)
		if targetPos != expected {
			t.Errorf("Expected %v, got %v", expected, targetPos)
		}
	})

	t.Run("src larger than target", func(t *testing.T) {
		src := Rect{900, 600}
		target := Rect{600, 400}
		srcPos := Vec{150, 100}
		expected := Vec{100, 66}
		targetPos := MapRectPos(src, target, srcPos)
		if targetPos != expected {
			t.Errorf("Expected %v, got %v", expected, targetPos)
		}
	})

	t.Run("src same size as target", func(t *testing.T) {
		src := Rect{600, 550}
		target := Rect{600, 550}
		srcPos := Vec{23, 45}
		expected := Vec{23, 45}
		targetPos := MapRectPos(src, target, srcPos)
		if targetPos != expected {
			t.Errorf("Expected %v, got %v", expected, targetPos)
		}
	})
}

func TestCenteringOffset(t *testing.T) {
	t.Run("Y offset is 0", func(t *testing.T) {
		inner := Rect{200, 400}
		outer := Rect{300, 400}
		expected := Vec{50, 0}
		offset := CenteringOffset(inner, outer)
		if offset != expected {
			t.Errorf("Expected %v, got %v", expected, offset)
		}
	})

	t.Run("Both X and Y positive", func(t *testing.T) {
		inner := Rect{200, 400}
		outer := Rect{300, 900}
		expected := Vec{50, 250}
		offset := CenteringOffset(inner, outer)
		if offset != expected {
			t.Errorf("Expected %v, got %v", expected, offset)
		}
	})

	t.Run("inner bigger than outer", func(t *testing.T) {
		inner := Rect{400, 900}
		outer := Rect{200, 300}
		expected := Vec{-100, -300}
		offset := CenteringOffset(inner, outer)
		if offset != expected {
			t.Errorf("Expected %v, got %v", expected, offset)
		}
	})

}
