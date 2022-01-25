package main

import (
	//"log"
	//"net/http"
	"os/exec"
)

const vlcPath = "/Applications/VLC.app/Contents/MacOS/VLC"

var runningVlc *exec.Cmd

func StopVlc() {
	if runningVlc != nil && runningVlc.Process != nil {
		runningVlc.Process.Kill()
	}
}

func main() {

	//mux := http.NewServeMux()
	//
	//mux.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
	//	streamQ, streamOk := r.URL.Query()["stream"]
	//	top, topOk := r.URL.Query()["top"]
	//	bottom, bottomOk := r.URL.Query()["bottom"]
	//	right, rightOk := r.URL.Query()["right"]
	//	left, leftOk := r.URL.Query()["left"]
	//
	//	if !streamOk {
	//		fmt.Fprintf(w, "Need to specify a stream.")
	//	}
	//	stream := string(streamQ[0])
	//
	//	crop := make(map[string]int)
	//	if topOk && len(top[0]) > 0 {
	//		t, err := strconv.Atoi(string(top[0]))
	//		if err == nil {
	//			crop["top"] = t
	//		}
	//	}
	//	if bottomOk && len(bottom[0]) > 0 {
	//		b, err := strconv.Atoi(string(bottom[0]))
	//		if err == nil {
	//			crop["bottom"] = b
	//		}
	//	}
	//	if leftOk && len(left[0]) > 0 {
	//		l, err := strconv.Atoi(string(left[0]))
	//		if err == nil {
	//			crop["left"] = l
	//		}
	//	}
	//	if rightOk && len(right[0]) > 0 {
	//		ri, err := strconv.Atoi(string(right[0]))
	//		if err == nil {
	//			crop["right"] = ri
	//		}
	//	}
	//
	//	vlcArgs := []string{
	//		vlc.CropFilterArg(crop),
	//		"--fullscreen",
	//	}
	//
	//	fmt.Print(vlcArgs)
	//	runningVlc = vlc.Play(vlcPath, stream, vlcArgs)
	//})
	//
	//mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
	//	StopVlc()
	//})
	//
	//log.Fatal(http.ListenAndServe(":8081", mux))
	//
	//cropArg := vlc.CropFilterArg(map[string]int{
	//	"top":    100,
	//	"bottom": 100,
	//	"left":   500,
	//	"right":  200,
	//})
	//vlcArgs := []string{"-vvv", cropArg, "--fullscreen"}
	//
	//fmt.Println(vlcArgs)
	//runningVlc = vlc.Play(vlc.OsVlcPath(),"/Users/miller/Documents/Code/wall-of-globes/terminator.mp4", vlcArgs)
	//
	//runningVlc.Wait()
}
