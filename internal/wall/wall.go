package wall

import (
	"encoding/json"
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
	if err != nil {
		return wl, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return wl, err
	}

	err = json.Unmarshal(bytes, &wl)
	if err != nil {
		return wl, err
	}

	return wl, nil
}
