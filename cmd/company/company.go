package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/geom"
	"github.com/millerpeterson/wall-of-globes/internal/media"
	"github.com/millerpeterson/wall-of-globes/internal/player"
	"github.com/millerpeterson/wall-of-globes/internal/tiling"
	"github.com/millerpeterson/wall-of-globes/internal/vlc"
	"github.com/millerpeterson/wall-of-globes/internal/wall"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type GlobePlayRequest struct {
	Address string
	Stream  string
	Args    player.Args
}

func playReqs(wallFile string, stream string, streamDim geom.Rect) []GlobePlayRequest {
	wl, err := wall.Load(wallFile)
	if err != nil {
		panic(err)
	}
	crops := tiling.Crops(streamDim, wl.Tln)
	var reqs []GlobePlayRequest
	for tlName, crop := range crops {
		reqs = append(reqs, GlobePlayRequest{
			Address: wl.ServerMap[tlName],
			Stream:  stream,
			Args:    crop,
		})
	}
	return reqs
}

func doReq(req GlobePlayRequest, ch chan<- string) {
	playUrl := fmt.Sprintf("http://%s/play", req.Address)
	ds, _ := json.Marshal(req.Args)
	log.Printf("Calling %v, data: %s", playUrl, ds)
	postData := url.Values{}
	postData.Set("file", req.Stream)
	postData.Set("top", strconv.Itoa(req.Args.Top))
	postData.Set("bottom", strconv.Itoa(req.Args.Bottom))
	postData.Set("left", strconv.Itoa(req.Args.Left))
	postData.Set("right", strconv.Itoa(req.Args.Right))
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Post(
		playUrl,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(postData.Encode()),
	)
	if err != nil {
		ch <- fmt.Sprintf("Req error: %v", err)
		return
	}
	ch <- fmt.Sprintf("%v -> %v", req.Address, resp.StatusCode)
}

func multicastStream(filePath string, multicastAddr string) *exec.Cmd {
	vlcArgs := []string{
		// "-vvv",
		"--sout",
		fmt.Sprintf("udp:%s", multicastAddr),
		"--loop",
	}
	return vlc.Play(filePath, vlcArgs)
}

func mediaDims(filePath string) (geom.Rect, error) {
	dims, err := media.MediaDims(filePath)
	if err != nil {
		log.Println("Failed to extract media dimensions; looking for STREAM_WIDTH / STREAM HEIGHT env vars.")
		dims.Width, err = strconv.Atoi(os.Getenv("STREAM_WIDTH"))
		if err != nil {
			return dims, err
		}
		dims.Height, err = strconv.Atoi(os.Getenv("STREAM_HEIGHT"))
		if err != nil {
			return dims, err
		}
	}
	return dims, nil
}

func main() {
	wallConfig := os.Args[1]
	stream := os.Args[2]
	streamDims, err := mediaDims(stream)
	if err != nil {
		log.Fatalf("Failed to determine input file dims: %v\n", err)
	}

	multicastAddr := "225.0.0.1"
	log.Printf("Starting multicast stream of %s to udp:%s\n", stream, multicastAddr)
	vlcMulticast := multicastStream(stream, multicastAddr)

	multicastStreamUrl := fmt.Sprintf("udp://@%s", multicastAddr)
	reqs := playReqs(wallConfig, multicastStreamUrl, streamDims)
	ch := make(chan string)
	for _, req := range reqs {
		go doReq(req, ch)
	}
	for range reqs {
		log.Println(<-ch)
	}

	err = vlcMulticast.Wait()
	if err != nil {
		panic(err)
	}
}
