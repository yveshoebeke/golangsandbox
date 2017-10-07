package main

import (
  "fmt"
  "readnav"
  "dandb"
)

func main() {
  waypoint := readnav.Readnav()
  var legcount int = 1

  for i := range waypoint.Locations {
    if i > 0 {
      fmt.Printf("Leg: %d From %s to %s:", legcount, waypoint.Locations[i-1].Name,waypoint.Locations[i].Name)
      distance, bearing := dandb.Dandb(waypoint.Locations[i-1].Lat,waypoint.Locations[i-1].Long,waypoint.Locations[i].Lat,waypoint.Locations[i].Long)
      fmt.Printf("\tdistance: %0.2fkm bearing: %0.1f (waypoint transition: %s -> %s)\n",distance/1000,bearing,waypoint.Locations[i-1].Type,waypoint.Locations[i].Type)
      legcount++
    }
  }
}
