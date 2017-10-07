package readnav

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"myconfig"
)

type Waypoint struct {
	Locations []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Lat string `json:"lat"`
		Long string `json:"long"`
	} `json:"locations"`
}

func Readnav() *Waypoint {
	config := myconfig.Getconfig()
	var buffer bytes.Buffer

	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(config.Waypoints.Navplan)
	var navplanfilename string = buffer.String()
	fmt.Println("Reading waypoint data from",navplanfilename)

	var Mywaypoint Waypoint
	navplanFile, err := os.Open(navplanfilename)
	defer navplanFile.Close()
	if err != nil {
		fmt.Println(Mywaypoint, err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(navplanFile)
	jsonParser.Decode(&Mywaypoint)

	return &Mywaypoint
}
