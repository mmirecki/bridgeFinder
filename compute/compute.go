package compute

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/dataquery"
	"github.com/mmirecki/bridgeFinder/known_uk_bridges"
	"github.com/mmirecki/bridgeFinder/lib"
)

func ComputeDebug(knownBridges map[int64]known_uk_bridges.KnownBridge) ([]*data.UnderWay, error) {

	//knownBridges := known_uk_bridges.GetKnownUKBridges()

	ways, err := ComputeForSouthampton(knownBridges)
	return ways, err

	//	printResults(ways, knownBridges)

}

func computeSquare(knownBridges map[int64]known_uk_bridges.KnownBridge, minLat, minLng, maxLat, maxLng float64) ([]*data.UnderWay, error) {

	inputDataSet, err := dataquery.NewDataSetForBounds(minLat, minLng, maxLat, maxLng)

	if err != nil {
		return nil, err
	}

	completeUnderWays := []*data.UnderWay{}

	inputBridgeCount := 0
	// For each bridge in the input data set
	for data, ok := inputDataSet.NextBridge(); ok; data, ok = inputDataSet.NextBridge() {
		inputBridgeCount++
		if len(data.PotentialUnderWays) == 0 {
			continue
		}

		underWays, err := lib.ExtractCrossRoads(data.Bridge, data.PotentialUnderWays)
		if err != nil {
			return nil, err
		}
		GetNeighbours(underWays, data.PotentialUnderWays)

		CheckKnownBridges(underWays, knownBridges)

		completeUnderWays = append(completeUnderWays, underWays...)
	}

	fmt.Printf("Input bridge count: %v\n", inputBridgeCount)
	return completeUnderWays, nil
}

func ComputeForSouthampton(knownBridges map[int64]known_uk_bridges.KnownBridge) ([]*data.UnderWay, error) {

	inputDataSet, err := dataquery.NewDataSet()
	if err != nil {
		return nil, err
	}

	completeUnderWays := []*data.UnderWay{}

	inputBridgeCount := 0
	// For each bridge in the input data set
	for data, ok := inputDataSet.NextBridge(); ok; data, ok = inputDataSet.NextBridge() {
		inputBridgeCount++
		if len(data.PotentialUnderWays) == 0 {
			continue
		}

		underWays, err := lib.ExtractCrossRoads(data.Bridge, data.PotentialUnderWays)
		if err != nil {
			return nil, err
		}
		GetNeighbours(underWays, data.PotentialUnderWays)

		CheckKnownBridges(underWays, knownBridges)

		completeUnderWays = append(completeUnderWays, underWays...)
	}

	fmt.Printf("Input bridge count: %v\n", inputBridgeCount)
	return completeUnderWays, nil
}
