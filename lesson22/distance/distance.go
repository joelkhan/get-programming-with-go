package main

import (
	"fmt"
	"math"
)

// world with a volumetric mean radius in kilometers
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

var (
	mercury = world{radius: 2439.7}
	venus   = world{radius: 6051.8}
	earth   = world{radius: 6371.0}
	mars    = world{radius: 3389.5}
	jupiter = world{radius: 69911.0}
	saturn  = world{radius: 58232.0}
	uranus  = world{radius: 25362.0}
	neptune = world{radius: 24622.0}
)

func main() {

	spirit := newLocation("spirit",
		coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation("opportunity",
		coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation("curiosity",
		coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insight := newLocation("insight",
		coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	locs := make([]location, 0, 4)
	locs = append(locs, spirit)
	locs = append(locs, opportunity)
	locs = append(locs, curiosity)
	locs = append(locs, insight)
	//fmt.Printf("%v\n", locs)
	//fmt.Println(len(locs))
	fmt.Println("on mars we have:")
	for i := 0; i < len(locs); i++ {
		fmt.Printf("%s -- ", locs[i].name)
	}
	fmt.Println()

	var farthest float64 = 0.0
	var farthestA, farthestB int
	var nearest float64 = 10000.0
	var nearestA, nearestB int
	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			dist := mars.distance(locs[i], locs[j])
			if dist > farthest {
				farthest = dist
				farthestA = i
				farthestB = j
			}
			if dist < nearest {
				nearest = dist
				nearestA = i
				nearestB = j
			}
			fmt.Printf("%15s -> %15s (%10.2f)\n", locs[i].name, locs[j].name, dist)
		}
	}
	fmt.Printf("farthest: %15s <--> %15s, %10.2f\n",
		locs[farthestA].name, locs[farthestB].name, farthest)
	fmt.Printf("nearest : %15s <--> %15s, %10.2f\n",
		locs[nearestA].name, locs[nearestB].name, nearest)

	aboutSolarSys()
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func aboutSolarSys() {
	london := newLocation("london",
		coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation("paris",
		coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	fmt.Printf("from London to Paris, about: %10.2f\n", earth.distance(london, paris))

	edmonton := newLocation("edmonton",
		coordinate{53, 32, 0, 'N'}, coordinate{113, 30, 0, 'W'})
	ottawa := newLocation("ottawa",
		coordinate{45, 25, 0, 'N'}, coordinate{75, 41, 0, 'W'})
	fmt.Printf("author's hometown to capital, about: %10.2f\n", earth.distance(edmonton, ottawa))

	sharp := newLocation("sharp",
		coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympus := newLocation("olympus",
		coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	fmt.Printf("Mount sharp to olympus, about: %10.2f\n", mars.distance(sharp, olympus))
}
