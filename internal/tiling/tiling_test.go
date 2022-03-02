package tiling

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"reflect"
	"testing"
)

func TestEnclosingRect(t *testing.T) {
	wall := Tiling{
		"one": Tile{
			Offset: geom.Vec{200, 300},
			Rect:   geom.Rect{500, 900},
		},
		"two": Tile{
			Offset: geom.Vec{300, 150},
			Rect:   geom.Rect{300, 1000},
		},
		"three": Tile{
			Offset: geom.Vec{0, 800},
			Rect:   geom.Rect{100, 1000},
		},
	}
	expected := geom.Rect{700, 1800}
	enclosing := EnclosingRect(wall)
	if enclosing != expected {
		t.Errorf("Expected %v, got %v", expected, enclosing)
	}
}

func TestCrops(t *testing.T) {
	// 3x3 grid
	wall := Tiling{
		// 0,0
		"one": Tile{
			Offset: geom.Vec{0, 0},
			Rect:   geom.Rect{30, 20},
		},
		// 1,0
		"two": Tile{
			Offset: geom.Vec{30, 0},
			Rect:   geom.Rect{30, 20},
		},
		// 2,0
		"three": Tile{
			Offset: geom.Vec{60, 0},
			Rect:   geom.Rect{30, 20},
		},
		// 0,1
		"four": Tile{
			Offset: geom.Vec{0, 20},
			Rect:   geom.Rect{30, 20},
		},
		// 1,1
		"five": Tile{
			Offset: geom.Vec{30, 20},
			Rect:   geom.Rect{30, 20},
		},
		// 2,1
		"six": Tile{
			Offset: geom.Vec{60, 20},
			Rect:   geom.Rect{30, 20},
		},
		// 0,2
		"seven": Tile{
			Offset: geom.Vec{0, 40},
			Rect:   geom.Rect{30, 20},
		},
		// 1,2
		"eight": Tile{
			Offset: geom.Vec{30, 40},
			Rect:   geom.Rect{30, 20},
		},
		// 2,2
		"nine": Tile{
			Offset: geom.Vec{60, 40},
			Rect:   geom.Rect{30, 20},
		},
	}
	srcVideo := geom.Rect{900, 600}
	crops := Crops(srcVideo, wall)
	expected := map[string]player.Args{
		"one": {
			Top:    0,
			Bottom: 400,
			Left:   0,
			Right:  600,
		},
		"two": {
			Top:    0,
			Bottom: 400,
			Left:   300,
			Right:  300,
		},
		"three": {
			Top:    0,
			Bottom: 400,
			Left:   600,
			Right:  0,
		},
		"four": {
			Top:    200,
			Bottom: 200,
			Left:   0,
			Right:  600,
		},
		"five": {
			Top:    200,
			Bottom: 200,
			Left:   300,
			Right:  300,
		},
		"six": {
			Top:    200,
			Bottom: 200,
			Left:   600,
			Right:  0,
		},
		"seven": {
			Top:    400,
			Bottom: 0,
			Left:   0,
			Right:  600,
		},
		"eight": {
			Top:    400,
			Bottom: 0,
			Left:   300,
			Right:  300,
		},
		"nine": {
			Top:    400,
			Bottom: 0,
			Left:   600,
			Right:  0,
		},
	}
	if !reflect.DeepEqual(crops, expected) {
		t.Errorf("Expected %v, got %v", expected, crops)
	}
}
