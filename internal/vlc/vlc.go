package vlc

import (
	"fmt"
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
		return []string{"--codec avcodec,none"}
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

func CropFilterArg(cropArgs PlayerArgs) string {
	cropArgStrings := map[string]string{
		"croptop":    strconv.Itoa(cropArgs.top),
		"cropbottom": strconv.Itoa(cropArgs.bottom),
		"cropleft":   strconv.Itoa(cropArgs.left),
		"cropright":  strconv.Itoa(cropArgs.right),
	}
	return VideoFilterArg("croppadd", cropArgStrings)
}

type PlayerArgs struct {
	top    int
	bottom int
	left   int
	right  int
}

type Player interface {
	play(file string, args PlayerArgs)
}

type AppPlayer struct {
	proc *exec.Cmd
}

func (p AppPlayer) play(file string, args PlayerArgs) {
	p.proc = Play(file, []string{CropFilterArg(args)})
}

type PlayCmd struct {
	file string
	args PlayerArgs
}

type PlayCmdLogger struct {
	log []PlayCmd
}

func (p PlayCmdLogger) play(file string, args PlayerArgs) {
	p.log = append(p.log, PlayCmd{file, args})
}

func Logger() *Player {
	var p Player
	p = PlayCmdLogger{}
	return &p
}
