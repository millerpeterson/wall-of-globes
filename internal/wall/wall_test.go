package wall

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/tiling"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	expected := Wall{
		Tln: tiling.Tiling{
			"one": tiling.Tile{
				Offset: geom.Vec{10, 10},
				Rect:   geom.Rect{300, 200},
			},
			"two": tiling.Tile{
				Offset: geom.Vec{50, 70},
				Rect:   geom.Rect{600, 400},
			},
			"three": tiling.Tile{
				Offset: geom.Vec{250, 270},
				Rect:   geom.Rect{900, 600},
			},
		},
		ServerMap: map[string]string{
			"one":   "192.168.0.1",
			"two":   "192.168.0.2",
			"three": "192.168.0.3",
		},
	}
	loaded, err := Load("test_wall.json")

	if err != nil {
		t.Errorf("Failed to load wall data: %v", err)
	}

	if !reflect.DeepEqual(expected, loaded) {
		t.Errorf("Expected %v, got %v", expected, loaded)
	}
}
