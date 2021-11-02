package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const vlcPath = "/Applications/VLC.app/Contents/MacOS/VLC"

func VlcPlay(file string, args []string) {
	cmdArgs := append([]string{file}, args...)
	vlcCmd := exec.Command(vlcPath, cmdArgs...)
	vlcCmd.Stderr = os.Stderr
	vlcCmd.Stdout = os.Stdout
	err := vlcCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	err = vlcCmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}

func VideoFilterVlcArgs(filterName string, filterArgs map[string]string) string {
	filterArgPairs := make([]string, len(filterArgs))
	argIndex := 0
	for argName := range filterArgs {
		filterArgPairs[argIndex] = fmt.Sprintf("%v=%v", argName, filterArgs[argName])
		argIndex++
	}
	return fmt.Sprintf("--video-filter=%v{%v}", filterName, strings.Join(filterArgPairs, ","))
}

func main() {
	VlcPlay("/Users/miller/Documents/Code/wall-of-globes/terminator.mp4", []string{})
}
