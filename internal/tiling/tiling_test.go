package tiling

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"reflect"
	"testing"
)

func TestEnclosingRect(t *testing.T) {
	wall := Tiling{
		Tile{
			Offset: geom.Vec{200, 300},
			Rect:   geom.Rect{500, 900},
		},
		Tile{
			Offset: geom.Vec{300, 150},
			Rect:   geom.Rect{300, 1000},
		},
		Tile{
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
		Tile{
			Offset: geom.Vec{0, 0},
			Rect:   geom.Rect{300, 200},
		},
		// 1,0
		Tile{
			Offset: geom.Vec{300, 0},
			Rect:   geom.Rect{300, 200},
		},
		// 2,0
		Tile{
			Offset: geom.Vec{600, 0},
			Rect:   geom.Rect{300, 200},
		},
		// 0,1
		Tile{
			Offset: geom.Vec{0, 200},
			Rect:   geom.Rect{300, 200},
		},
		// 1,1
		Tile{
			Offset: geom.Vec{300, 200},
			Rect:   geom.Rect{300, 200},
		},
		// 2,1
		Tile{
			Offset: geom.Vec{600, 200},
			Rect:   geom.Rect{300, 200},
		},
		// 0,2
		Tile{
			Offset: geom.Vec{0, 400},
			Rect:   geom.Rect{300, 200},
		},
		// 1,2
		Tile{
			Offset: geom.Vec{300, 400},
			Rect:   geom.Rect{300, 200},
		},
		// 2,2
		Tile{
			Offset: geom.Vec{600, 400},
			Rect:   geom.Rect{300, 200},
		},
	}
	srcVideo := geom.Rect{900, 600}
	crops := Crops(srcVideo, wall)
	expected := []player.Args{
		{
			Top:    0,
			Bottom: 400,
			Left:   0,
			Right:  600,
		},
		{
			Top:    0,
			Bottom: 400,
			Left:   300,
			Right:  300,
		},
		{
			Top:    0,
			Bottom: 400,
			Left:   600,
			Right:  0,
		},
		{
			Top:    200,
			Bottom: 200,
			Left:   0,
			Right:  600,
		},
		{
			Top:    200,
			Bottom: 200,
			Left:   300,
			Right:  300,
		},
		{
			Top:    200,
			Bottom: 200,
			Left:   600,
			Right:  0,
		},
		{
			Top:    400,
			Bottom: 0,
			Left:   0,
			Right:  600,
		},
		{
			Top:    400,
			Bottom: 0,
			Left:   300,
			Right:  300,
		},
		{
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
