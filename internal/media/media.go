package media

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"os/exec"
	"regexp"
	"strconv"
)

func MediaDims(mediaPath string) (geom.Rect, error) {
	dims := geom.Rect{}
	ffprobePath, err := exec.LookPath("ffprobe")
	if err != nil {
		return dims, err
	}
	ffprobeArgs := []string{
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=width,height",
		"-of",
		"csv=s=x:p=0",
		mediaPath,
	}
	dimsOut, err := exec.Command(ffprobePath, ffprobeArgs...).Output()
	dimsStr := string(dimsOut)
	dimsRegexp := regexp.MustCompile("^(?P<width>[0-9]+)x(?P<height>[0-9]+)")
	match := dimsRegexp.FindStringSubmatch(dimsStr)
	if len(match) == 0 {
		return dims, fmt.Errorf("Can't parse ffprobe output: %s", dimsStr)
	}
	dims.Width, err = strconv.Atoi(match[1])
	if err != nil {
		return dims, err
	}
	dims.Height, err = strconv.Atoi(match[2])
	if err != nil {
		return dims, err
	}
	return dims, nil
}
