package wall

import (
	"encoding/json"
	"fmt"
	"github.com/millerpeterson/wall-of-globes/internal/tiling"
	"io/ioutil"
	"os"
)

type TileServerMap map[string]string

type Wall struct {
	Tln tiling.Tiling `json:"tiles"`
	// Map of tile name -> server address
	ServerMap TileServerMap `json:"server_map"`
}

func Load(fp string) (Wall, error) {
	wl := Wall{}
	jsonFile, err := os.Open(fp)
	fmt.Println("1")
	if err != nil {
		return wl, err
	}
	defer func(jsonFile *os.File) {
		fmt.Println("2")
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	fmt.Println("3")
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return wl, err
	}

	fmt.Println("4")
	err = json.Unmarshal(bytes, &wl)
	if err != nil {
		return wl, err
	}

	fmt.Println("5")
	return wl, nil
}
