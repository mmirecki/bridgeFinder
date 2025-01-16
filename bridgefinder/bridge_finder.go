package main

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/compute"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/known_uk_bridges"
	"github.com/mmirecki/bridgeFinder/utils"
)

func main() {

	knownBridges := known_uk_bridges.GetKnownUKBridges()

	//ways := computeForArea(knownBridges, 50.98, -1.5, 51.0, -1.48)
	//ways := computeForArea(knownBridges, 50.98, -1.5, 51.0, -1.48)
	// (50.89, -1.34, 50.96, -1.52); - southampton
	//ways, err := compute.ComputeArea(knownBridges, data.LatLng{Lng: -1.5, Lat: 50.8}, data.LatLng{Lng: -1.48, Lat: 51.0})

	//if true {

	useCache := true

	ways, err := compute.ComputeArea(knownBridges, data.LatLng{Lng: utils.UK_MIN_LNG, Lat: utils.UK_MIN_LAT}, data.LatLng{Lng: utils.UK_MAX_LNG, Lat: utils.UK_MAX_LAT}, useCache)
	//ways, err := compute.ComputeArea(knownBridges, data.LatLng{Lng: -0.0, Lat: 51.0}, data.LatLng{Lng: 1.0, Lat: 52.0}, useCache)
	//}
	if err != nil {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&\n")
		fmt.Printf("Error: %v\n", err)
	}
	//ways := computeDebug(knownBridges)

	fmt.Printf("Count ways: %d\n", len(ways))

	//printResults(ways)

}
