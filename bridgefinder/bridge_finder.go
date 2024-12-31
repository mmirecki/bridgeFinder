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
	ways, err := compute.ComputeArea(knownBridges, data.LatLng{Lng: utils.UK_MIN_LNG, Lat: utils.UK_MIN_LAT}, data.LatLng{Lng: utils.UK_MAX_LNG, Lat: utils.UK_MAX_LAT})
	if err != nil {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&\n")
		fmt.Printf("Error: %v\n", err)
	}
	//ways := computeDebug(knownBridges)

	fmt.Printf("Count ways: %d\n", len(ways))

	//printResults(ways)

}
