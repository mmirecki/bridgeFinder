package reporting

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/lib"
)

func printResults(completeUnderWays []*data.UnderWay) {

	/*
		for _, underWay := range completeUnderWays {
			knownBridge, ok := knownBridges[underWay.Way.Id]
			if ok {
				underWay.IsExactKnownBridge = true
				underWay.KnownBridge = knownBridge
			}
		}

	*/

	missingCount := 0
	knownCount := 0
	hasNeighbourCount := 0

	for _, cr := range completeUnderWays {
		cr.CameraPositions = lib.GetCameraPositionsForWay(*cr)
	}

	fmt.Printf("===== MISSING BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && !underWay.HasNeighbouringKnownBridge {
			missingCount++
			printUnderway(underWay, false)
		}
	}

	fmt.Printf("\n\n\n===== KNOWN NEIGHBOUR BRIDGES ===========\n")

	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && underWay.HasNeighbouringKnownBridge {
			hasNeighbourCount++
			printUnderway(underWay, true)
		}
	}

	fmt.Printf("\n\n\n===== KNOWN BRIDGES ===========\n")

	for _, underWay := range completeUnderWays {
		if underWay.IsExactKnownBridge {
			knownCount++
			printUnderway(underWay, false)
		}
	}

	fmt.Printf("Underways len: %v\n", len(completeUnderWays))
	fmt.Printf("Known count: %v\n", knownCount)
	fmt.Printf("Neighbor known count: %v\n", hasNeighbourCount)
	fmt.Printf("Missing count: %v\n", missingCount)

}

func printUnderway(underWay *data.UnderWay, printNeighbours bool) {
	fmt.Printf("--------\n")
	fmt.Printf("  BRIDGE: %d %s   %s\n", underWay.Overhead.Id, underWay.Overhead.Tags["name"], fmt.Sprintf("https://www.openstreetmap.org/way/%d", underWay.Overhead.Id))
	fmt.Printf("     Underway %d \"%s\" Height:\"%s\"  %s\n", underWay.Way.Id, underWay.Way.Tags["name"], underWay.MaxHeight, fmt.Sprintf("https://www.openstreetmap.org/way/%d", underWay.Way.Id))
	if printNeighbours {
		for _, neighbour := range underWay.KnownNeighbours {
			fmt.Printf("          KNOWN Neighbour %d \"%s\" Height:\"%s\"  %s\n", neighbour.Id, neighbour.Tags["name"], neighbour.MaxHeight, fmt.Sprintf("https://www.openstreetmap.org/way/%d", neighbour.Id))
		}
	}

	fmt.Printf("           StreetView: ")
	for _, position := range underWay.CameraPositions {
		fmt.Printf("              %+v %s  \n", position.Position, position.ImageLink)
	}
}
