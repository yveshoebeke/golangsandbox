// ssavnav.go
// Calcuates distance between 2 geo-locations
// This was its own package to be utilized in the overall nav functionality
package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
  "crypto/md5"
  "encoding/hex"
  "time"
  "runtime"
  )

// constants needed
const VERSION = "1.0b"          // version
const DEG2RAD = math.Pi / 180   // degrees to rads conversion
const ARGSNEEDED = 5            // CLI arguments needed (+1 for exec name)
const EARTHRADIUS = 6378100     // Earth radius in meters (aprox)
const M2KM = 0.001              // meters to kilometers conversion
const M2SM = 0.0006214          // meters to statute miles
const M2NM = 0.0005399          // meters to nautical miles

// haversin function 
// (look it up - Wikipedia got a good explanation (curvature of the Earth stuff)
func hsin(theta float64) float64 {
  return math.Pow(math.Sin(theta/2), 2)
}

// get some data 
// (originally from GPS device output)
// (could be adapted coming from menu or html frontend)
func getdata() (float64, float64, float64, float64) {
  lat1, _ := strconv.ParseFloat(os.Args[1], 64)
  lon1, _ := strconv.ParseFloat(os.Args[2], 64)
  lat2, _ := strconv.ParseFloat(os.Args[3], 64)
  lon2, _ := strconv.ParseFloat(os.Args[4], 64)

  return lat1*DEG2RAD, lon1*DEG2RAD, lat2*DEG2RAD, lon2*DEG2RAD
}

// calc MD5 hash (not really needed in this stand alone version
func getMD5(text string) string {
  hasher := md5.New()
  hasher.Write([]byte(text))
  return hex.EncodeToString(hasher.Sum(nil))
}

// show user results (could be adapted to file or html output)
func showresults(distance, bearing float64) {
  fmt.Printf("\nDistance = %0.2f meters. Bearing = %0.1f degrees True North.\n", distance, bearing)
  fmt.Printf("Alternative distances: %0.2fkm - %0.2fsm - %0.2fnm.\n", distance*M2KM, distance*M2SM, distance*M2NM)

  fmt.Printf("\n[exec: %s v%s MD5:%s on %s at %s]\n", os.Args[0], VERSION, getMD5(os.Args[0]), runtime.GOOS, time.Now())
  return
}

// the geo-location structures (properties)
type Geopoint struct {
  lat1, lon1, lat2, lon2, distance, bearing, h, x, y float64
}

// g method to get distance subject to the structure
func (g *Geopoint) Distance() float64 {
  g.h = hsin(g.lat2 - g.lat1) + math.Cos(g.lat1) * math.Cos(g.lat2) * hsin(g.lon2 - g.lon1)
  g.distance = 2 * EARTHRADIUS * math.Asin(math.Sqrt(g.h))
  
  return g.distance
 }

// g method to get bearing subject to the structure
func (g *Geopoint) Bearing() float64 {
  g.x = math.Cos(g.lat2) * math.Sin(g.lon2 - g.lon1)
  g.y = math.Cos(g.lat1) * math.Sin(g.lat2) - math.Sin(g.lat1) * math.Cos(g.lat2) * math.Cos(g.lon2 - g.lon1)
  g.bearing = math.Atan2(g.x, g.y)
 
  if g.bearing < 0 {
    g.bearing = math.Abs(g.bearing) + (90 * DEG2RAD)
  }
  
  return g.bearing / DEG2RAD
 }


 //putting it all together
func main() {
  var lat1, lon1, lat2, lon2, distance, bearing, h, x, y float64 = 0, 0, 0, 0, 0, 0, 0, 0, 0
  var latNS, lonEW string = "N", "E"

  if len(os.Args) < ARGSNEEDED {
    println("Not enough arguments - need 4 (latitude1 longitude1 latitude2 longitude2)");
    os.Exit(0)
  } else {
    lat1, lon1, lat2, lon2 = getdata()

    fmt.Println("\nDistance and Bearing between:")
    
    // Some cosmetics to replace signage with real-world verbage
    if lat1 < 0 {
      latNS = "S"
    }

    if lon1 < 0 {
      lonEW = "W"
    }

    fmt.Println("Origin - ", math.Abs(lat1), latNS, "lat - ", math.Abs(lon1), lonEW,  "long.")

    if lat2 < 0 {
      latNS = "S"
    } else {
      latNS = "N"
    }

    if lon2 < 0 {
      lonEW = "W"
    } else {
      lonEW = "E"
    }

    fmt.Println("Target - ", math.Abs(lat2), latNS, "lat - ", math.Abs(lon2), lonEW, "long.")
  }
 
  // intitialize the structure from input data etc
  g := Geopoint{lat1, lon1, lat2, lon2, distance, bearing, h, x, y}

  // get distance and bearing from methods
  distance = g.Distance()
  bearing = g.Bearing()

  // show results
  showresults(distance, bearing)


}

// End of file ssavnav.go
