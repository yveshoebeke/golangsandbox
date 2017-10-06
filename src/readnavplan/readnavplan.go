package main

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
)


type Config struct {
	Waypoints struct {
		Navplan string `json:"navplan"`
	} `json:"waypoints"`
}

type Waypoint struct {
	Locations []struct { 
		Name string `json:"name"`
		Type string `json:"type"`
		Lat string `json:"lat"`
		Long string `json:"long"`
	} `json:"locations"`
}

func LoadConfig(file string) (Config, error) {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
		return config, err
	}
    
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config, err
}

func main() {
	var buffer bytes.Buffer
	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString("/config/config.json")
	var configfilename string = buffer.String()


	fmt.Println("Running...")
	config, _ := LoadConfig(configfilename)

	buffer.Reset()
	buffer.WriteString(os.Getenv("GOPATH"))
	buffer.WriteString(config.Waypoints.Navplan)
	var navplanfilename string = buffer.String()
	fmt.Println("Reading waypoint data from",navplanfilename)

	// cycle thru the navplan waypoints
	var waypoint Waypoint
	navplanFile, err := os.Open(navplanfilename)
	defer navplanFile.Close()
	if err != nil {
		fmt.Println(waypoint, err.Error())
		os.Exit(1)
	}
	
	jsonParser := json.NewDecoder(navplanFile)
	jsonParser.Decode(&waypoint)
	
	fmt.Println("Waypoints:")
	fmt.Println("Name\t\tType\t\tLatitude\tLongitude")
	for i := range waypoint.Locations {
		fmt.Printf("%s\t%s\t\t%s\t\t%s\n",waypoint.Locations[i].Name,waypoint.Locations[i].Type,waypoint.Locations[i].Lat,waypoint.Locations[i].Long)
	}

}

