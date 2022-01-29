package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// location with a latitude, longitude.
type location struct {
	name      string
	lat, long float64
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(name string, lat, long coordinate) location {
	return location{name, lat.decimal(), long.decimal()}
}

func (loc location) desc() string {
	return fmt.Sprintf("name: %s (lat: %.2f, long: %.2f)", loc.name, loc.lat, loc.long)
}

type gps struct {
	world
	locCurrent location
	locTarget  location
}

func (g gps) distance() float64 {
	return g.world.distance(g.locCurrent, g.locTarget)
}

func (g gps) message() string {
	return fmt.Sprintf("To %s, about: %.2f km", g.locTarget.desc(), g.distance())
}

func newGPS(w world, loc1, loc2 location) gps {
	return gps{w, loc1, loc2}
}

type rover struct {
	name string
	gps
}

// func (r rover) message() string {

// }

func newRover(name string, g gps) rover {
	return rover{name, g}
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{"bradbury", -4.5895, 137.4417}
	elysium := location{"elysium", 4.5, 135.9}
	curiosity := newRover("curiosity", newGPS(mars, bradbury, elysium))
	//fmt.Println(curiosity)
	fmt.Println(curiosity.message())
}
