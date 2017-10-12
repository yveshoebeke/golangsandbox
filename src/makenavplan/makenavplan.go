// makenavplan
// reads the waypoints out of a json file and calculates distance and bearing for each leg.
package main

import (
	"dandb" // calculate distance and bearing
	"fmt"
	"logusage" // db access to log/display session
	"myconfig" // get this app's configuration
	"os"
	"readnav" // get waypoints
)

func main() {
	config := myconfig.Getconfig()
	waypoint := readnav.Readnav()
	// display application name and version
	fmt.Printf("\n%s (Version %s)\n\n", config.Application.Name, config.Application.Version)
	// display this user' previous access sessions to this application
	fmt.Printf("Previous access by %s:\n", os.Getenv("USER"))
	logusage.Logaction("showit", "")
	fmt.Println()
	// iterate through the waypoints and display each segment with it's pertinent data on the console.
	for i := range waypoint.Locations {
		if i > 0 {
			fmt.Printf("Leg: %d From %s to %s:", i, waypoint.Locations[i-1].Name, waypoint.Locations[i].Name)
			distance, bearing := dandb.Dandb(waypoint.Locations[i-1].Lat, waypoint.Locations[i-1].Long, waypoint.Locations[i].Lat, waypoint.Locations[i].Long)
			fmt.Printf("\tdistance: %0.2fkm bearing: %0.1f (waypoint transition: %s -> %s)\n", distance*config.Conversions.M2km, bearing, waypoint.Locations[i-1].Type, waypoint.Locations[i].Type)
		}
	}
	// log this session
	logusage.Logaction("logit", os.Getenv("USER"))
	fmt.Println("\nDone")
}
