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
	Application struct {
		Name string `json:"name"`
		Version string `json:"version"`
		Author string `json:"author"`
	}
	Navdata struct {
	  Navplan string `json:"navplan"`
	} `json:"waypoints"`
  Database struct {
    Host string `json:"host"`
    Port string `json:"port"`
    User string `json:"user"`
    Password string `json:"password"`
    Database string `json:"database"`
  } `json:"database"`
  Conversions struct {
    M2km float64 `json:"m2km"`
    M2sm float64 `json:"M2sm"`
    M2nm float64 `json:"M2nm"`
  } `json:"conversions"`
  Constants struct {
    Earthradius float64 `json:"earthradius"`
  }
}

func Getconfig() *Config {
	var Myconfig Config
	var buffer bytes.Buffer

        buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(CONFIGPATH)
	var configfilename string = buffer.String()

	configFile, err := os.Open(configfilename)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Myconfig); err != nil {
		fmt.Println("Configuration decoding error:", err.Error())
		os.Exit(1)
	}
	return &Myconfig
}
