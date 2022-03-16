package main

import (
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/wall"
	"log"
	"os"
	"os/exec"
	"strings"
)

const sshUser string = "pi"
const binary string = "bin/globe-pi"

func rebootPi(piIp string) *exec.Cmd {
	sshArgs := []string{
		fmt.Sprintf("%s@%s", sshUser, piIp),
		fmt.Sprintf("sudo shutdown -r now"),
	}
	cmd := exec.Command("ssh", sshArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return cmd
}

func wallIps(wl wall.Wall) []string {
	ips := []string{}
	for _, addr := range wl.ServerMap {
		addrParts := strings.Split(addr, ":")
		if len(addrParts) == 2 {
			ips = append(ips, addrParts[0])
		} else {
			log.Printf("Failed to parse server %s", addr)
		}
	}
	return ips
}

func rebootWall(wl wall.Wall) {
	for _, ip := range wallIps(wl) {
		log.Printf("Rebooting %s", ip)
		rebootPi(ip)
	}
}

func syncGlobeBinary(binary string, piAddr string) *exec.Cmd {
	scpArgs := []string{
		binary,
		fmt.Sprintf("%s@%s:~/globe", sshUser, piAddr),
	}
	cmd := exec.Command("scp", scpArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return cmd
}

func syncWallGlobes(wl wall.Wall) {
	for _, ip := range wallIps(wl) {
		log.Printf("Syncing to %s", ip)
		syncGlobeBinary(binary, ip).Wait()
	}
}

func main() {
	action := os.Args[1]

	wallFile := os.Args[2]
	wl, err := wall.Load(wallFile)
	if err != nil {
		log.Fatalf("Failed to load wall JSON: %v", err)
	}

	switch action {
	case "reboot":
		rebootWall(wl)
	case "sync":
		syncWallGlobes(wl)
	}
}
