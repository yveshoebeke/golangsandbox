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
	// construct path to nav data
	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(config.Navdata.Navplan)
	navplanfilename  = buffer.String()
	// try and open nav data file
	navplanFile, err := os.Open(navplanfilename)
	defer navplanFile.Close()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	// parse nav json data
	jsonParser := json.NewDecoder(navplanFile)
	if err = jsonParser.Decode(&Mywaypoint); err != nil {
		fmt.Println("Navigation data decoding error:", err.Error())
		os.Exit(1)
	}

	return &Mywaypoint
}
