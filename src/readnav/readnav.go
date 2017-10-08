// readnav
// reads the nav waypoints out of a json file.
package readnav

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"myconfig"
)

// waypoint construct
type Waypoint struct {
	Locations []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Lat string `json:"lat"`
		Long string `json:"long"`
	} `json:"locations"`
}

// reads all the waypoint data and 'inits' the structure.
func Readnav() *Waypoint {
	config := myconfig.Getconfig()
	var Mywaypoint Waypoint
	var buffer bytes.Buffer
	var navplanfilename string

	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(config.Waypoints.Navplan)
	navplanfilename  = buffer.String()
	fmt.Println("Reading waypoint data from",navplanfilename)

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
