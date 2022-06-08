package media

import (
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"os"
	"reflect"
	"testing"
)

func TestMediaDims(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping test in CI (until we configure ffprobe for CI)")
	}
	expectedDims := geom.Rect{1280, 720}
	probedDims, err := MediaDims("Beam Sequence 900, Quantized Space.mp4")
	if err != nil {
		t.Fatalf("Failed to get media dims: %v", err)
	}
	if !reflect.DeepEqual(expectedDims, probedDims) {
		t.Errorf("Expected %v, got %v", expectedDims, probedDims)
	}
}
