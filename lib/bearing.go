package lib

import (
	"github.com/mmirecki/bridgeFinder/data"
	"math"
)

// CalculateBearing calculates the bearing from start to end coordinates
// func CalculateBearing(startLat, startLon, endLat, endLon float64) float64 {
func CalculateBearing(segment data.Segment) float64 {
	// Convert degrees to radians
	startLatRad := degreesToRadians(segment.Start.Lat)
	startLonRad := degreesToRadians(segment.Start.Lng)
	endLatRad := degreesToRadians(segment.End.Lat)
	endLonRad := degreesToRadians(segment.End.Lng)

	// Calculate the difference in longitudes
	deltaLon := endLonRad - startLonRad

	// Calculate the bearing
	x := math.Sin(deltaLon) * math.Cos(endLatRad)
	y := math.Cos(startLatRad)*math.Sin(endLatRad) - math.Sin(startLatRad)*math.Cos(endLatRad)*math.Cos(deltaLon)
	bearing := math.Atan2(x, y)

	// Convert bearing from radians to degrees
	bearing = radiansToDegrees(bearing)

	// Normalize the bearing to 0-360 degrees
	if bearing < 0 {
		bearing += 360
	}

	return bearing
}

// degreesToRadians converts degrees to radians
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// radiansToDegrees converts radians to degrees
func radiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
