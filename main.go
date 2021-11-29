package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
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

func CropFilterArg(cropArgs map[string]int) string {
	cropArgNames := map[string]string{
		"top":    "croptop",
		"bottom": "cropbottom",
		"left":   "cropleft",
		"right":  "cropright",
	}
	cropArgsStrings := make(map[string]string, len(cropArgs))
	for argName := range cropArgs {
		cropArgsStrings[cropArgNames[argName]] = strconv.Itoa(cropArgs[argName])
	}
	return VideoFilterArg("croppadd", cropArgsStrings)
}

func main() {
	cropArg := CropFilterArg(map[string]int{
		"top":    100,
		"bottom": 100,
		"left":   500,
		"right":  200,
	})
	vlcArgs := []string{cropArg}

	fmt.Print(vlcArgs)
	VlcPlay("/Users/miller/Documents/Code/wall-of-globes/terminator.mp4", vlcArgs)
}
