package media

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"reflect"
	"testing"
)

func TestMediaDims(t *testing.T) {
	expectedDims := geom.Rect{1280, 720}
	probedDims, err := MediaDims("Beam Sequence 900, Quantized Space.mp4")
	if err != nil {
		t.Fatalf("Failed to get media dims: %v", err)
	}
	if !reflect.DeepEqual(expectedDims, probedDims) {
		t.Errorf("Expected %v, got %v", expectedDims, probedDims)
	}
}
