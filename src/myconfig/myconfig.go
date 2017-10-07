// myconfig
// reads to configuration file
package myconfig

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
)

const CONFIGPATH = "/config/config.json"

type Config struct {
	Waypoints struct {
		Navplan string `json:"navplan"`
	} `json:"waypoints"`
}

func Getconfig() *Config {
	var config Config
	var buffer bytes.Buffer
	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(CONFIGPATH)
	var configfilename string = buffer.String()

	configFile, err := os.Open(configfilename)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Error opening", configfilename, err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}
