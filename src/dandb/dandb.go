// dandb
// Calcuates distance and bearing between 2 geo-locations
package dandb

import (
  "math"
  "strconv"
)

// constants needed
const DEG2RAD = math.Pi / 180   // degrees to rads conversion
const EARTHRADIUS = 6378100     // Earth radius in meters (aprox)

// haversin function
// (look it up - Wikipedia got a good explanation (curvature of the Earth stuff)
func hsin(theta float64) float64 {
  return math.Pow(math.Sin(theta/2), 2)
}

func getD2Rdata(slat1, slon1, slat2, slon2 string) (float64, float64, float64, float64) {
  lat1, _ := strconv.ParseFloat(slat1, 64)
  lon1, _ := strconv.ParseFloat(slon1, 64)
  lat2, _ := strconv.ParseFloat(slat2, 64)
  lon2, _ := strconv.ParseFloat(slon2, 64)

  return lat1*DEG2RAD, lon1*DEG2RAD, lat2*DEG2RAD, lon2*DEG2RAD
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

//entry point
func Dandb(slat1, slon1, slat2, slon2 string) (float64, float64) {
  var lat1, lon1, lat2, lon2, distance, bearing, h, x, y float64 = 0, 0, 0, 0, 0, 0, 0, 0, 0
  lat1, lon1, lat2, lon2 = getD2Rdata(slat1, slon1, slat2, slon2)

  // intitialize the structure from input data etc
  g := Geopoint{lat1, lon1, lat2, lon2, distance, bearing, h, x, y}

  // get distance and bearing from methods
  distance = g.Distance()
  bearing = g.Bearing()

  return distance, bearing
}
