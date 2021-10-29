package main

import (
	"os"
	"os/exec"
)

const vlcPath = "/Applications/VLC.app/Contents/MacOS/VLC"

func vlcPlay(file string, args []string) {
	cmdArgs := append([]string{file}, args...)
	vlcCmd := exec.Command(vlcPath, cmdArgs...)
	vlcCmd.Stderr = os.Stderr
	vlcCmd.Stdout = os.Stdout
	err := vlcCmd.Run()
	if err != nil {
		panic(err)
	}
	err = vlcCmd.Wait()
	if err != nil {
		panic(err)
	}
}

func main() {
	vlcPlay("/Users/miller/Documents/Code/wall-of-globes/terminator.mp4", []string{})
}
