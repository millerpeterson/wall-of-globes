package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

const vlcPath = "/Applications/VLC.app/Contents/MacOS/VLC"

func VlcPlay(file string, args []string) *exec.Cmd {
	cmdArgs := append([]string{file}, args...)
	vlcCmd := exec.Command(vlcPath, cmdArgs...)
	vlcCmd.Stderr = os.Stderr
	vlcCmd.Stdout = os.Stdout
	err := vlcCmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	return vlcCmd
	//err = vlcCmd.Wait()
	//if err != nil {
	//	fmt.Println(err)
	//}
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

var vlc *exec.Cmd

func StopVlc() {
	if vlc != nil && vlc.Process != nil {
		vlc.Process.Kill()
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		streamQ, streamOk := r.URL.Query()["stream"]
		top, topOk := r.URL.Query()["top"]
		bottom, bottomOk := r.URL.Query()["bottom"]
		right, rightOk := r.URL.Query()["right"]
		left, leftOk := r.URL.Query()["left"]

		if !streamOk {
			fmt.Fprintf(w, "Need to specify a stream.")
		}
		stream := string(streamQ[0])

		crop := make(map[string]int)
		if topOk && len(top[0]) > 0 {
			t, err := strconv.Atoi(string(top[0]))
			if err == nil {
				crop["top"] = t
			}
		}
		if bottomOk && len(bottom[0]) > 0 {
			b, err := strconv.Atoi(string(bottom[0]))
			if err == nil {
				crop["bottom"] = b
			}
		}
		if leftOk && len(left[0]) > 0 {
			l, err := strconv.Atoi(string(left[0]))
			if err == nil {
				crop["left"] = l
			}
		}
		if rightOk && len(right[0]) > 0 {
			ri, err := strconv.Atoi(string(right[0]))
			if err == nil {
				crop["right"] = ri
			}
		}

		vlcArgs := []string{
			CropFilterArg(crop),
			"--fullscreen",
		}

		fmt.Print(vlcArgs)
		vlc = VlcPlay(stream, vlcArgs)
	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		StopVlc()
	})

	log.Fatal(http.ListenAndServe(":8081", mux))

	//cropArg := CropFilterArg(map[string]int{
	//	"top":    100,
	//	"bottom": 100,
	//	"left":   500,
	//	"right":  200,
	//})
	//vlcArgs := []string{cropArg}
	//
	//fmt.Print(vlcArgs)
	//VlcPlay("/Users/miller/Documents/Code/wall-of-globes/terminator.mp4", vlcArgs)
}
