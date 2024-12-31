package utils

import "github.com/mmirecki/bridgeFinder/data"

const (
	UK_MIN_LAT = 49.8

	UK_MAX_LAT = 60.2
	UK_MIN_LNG = -8.6
	UK_MAX_LNG = 1.8

	UK_LNG_INCREMENT = 0.1
	UK_LAT_INCREMENT = 0.1
)

func IsLatLngInUk(lng data.LatLng) bool {

	if lng.Lng < UK_MIN_LNG || lng.Lng > UK_MAX_LNG {
		return false
	}
	if lng.Lat < UK_MIN_LAT || lng.Lat > UK_MAX_LAT {
		return false
	}
	// y=0.5953x+49.9966 - cut of that little bit of France in the south west of the rectangle
	if lng.Lat < 0.5953*lng.Lng+49.9966 {
		return false
	}
	return true
}
