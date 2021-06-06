package main

import (
	"fmt"
)

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// Stirng formats a DMS coordinate.
func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// locaiton with a latitude, longitude in decimal degrees.
type location struct {
	lat, long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0, 'E'},
	}

	fmt.Println("Elysium Planitia is at", elysium)
}
