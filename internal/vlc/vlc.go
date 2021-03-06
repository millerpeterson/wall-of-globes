package vlc

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func OsVlcPath() string {
	if runtime.GOOS == "darwin" {
		// NOTE: VLC does not seem to start properly on OSX if you are attempting
		// to use a simlink.
		return "/Applications/VLC.app/Contents/MacOS/VLC"
	} else {
		return "vlc"
	}
}

func OSVlcArgs() []string {
	if runtime.GOOS == "linux" {
		// The crop filter seems not to work on the Pi with hardware decoding.
		return []string{"--codec", "avcodec,none"}
	} else {
		return []string{}
	}
}

func Play(file string, args []string) *exec.Cmd {
	cmdArgs := append([]string{file}, args...)
	cmdArgs = append(cmdArgs, OSVlcArgs()...)
	vlcCmd := exec.Command(OsVlcPath(), cmdArgs...)
	vlcCmd.Stderr = os.Stderr
	vlcCmd.Stdout = os.Stdout
	err := vlcCmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	return vlcCmd
}

func VideoFilterArg(filterName string, filterArgs map[string]string) string {
	sortedArgs := make([]string, 0)
	for argName := range filterArgs {
		sortedArgs = append(sortedArgs, argName)
	}
	sort.Strings(sortedArgs)
	filterArgPairs := make([]string, len(filterArgs))
	for i, argName := range sortedArgs {
		filterArgPairs[i] = fmt.Sprintf("%v=%v", argName, filterArgs[argName])
	}
	return fmt.Sprintf("--video-filter=%v{%v}", filterName, strings.Join(filterArgPairs, ","))
}

func CropFilterArg(cropArgs player.Args) string {
	cropArgStrings := map[string]string{
		"croptop":    strconv.Itoa(cropArgs.Top),
		"cropbottom": strconv.Itoa(cropArgs.Bottom),
		"cropleft":   strconv.Itoa(cropArgs.Left),
		"cropright":  strconv.Itoa(cropArgs.Right),
	}
	return VideoFilterArg("croppadd", cropArgStrings)
}

type Player struct {
	proc *exec.Cmd
}

func (p *Player) Play(file string, args player.Args) {
	p.Stop()
	p.proc = Play(file, []string{"-vvv", "--fullscreen", CropFilterArg(args)})
}

func (p *Player) Stop() {
	if p.proc == nil {
		return
	}

	err := p.proc.Process.Kill()
	if err != nil {
		log.Printf("Warning: Failed to stop VLC process: %v", err)
	}
	p.proc = nil
}
