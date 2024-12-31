package compute

import (
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/known_uk_bridges"
)

func GetNeighbours(underways []*data.UnderWay, allWays []data.Way) {
	for _, underWay := range underways {
		neighbours := []*data.Way{}

		for _, node := range []data.Node{underWay.Way.Nodes[0], underWay.Way.Nodes[len(underWay.Way.Nodes)-1]} {
			for _, way := range allWays {
				if underWay.Way.Id == way.Id {
					continue
				}
				if way.Nodes[0] == node || way.Nodes[len(way.Nodes)-1] == node {
					neighbours = append(neighbours, &way)
				}
			}
		}
		underWay.Neighbours = neighbours
	}
}

func CheckKnownBridges(underways []*data.UnderWay, knownBridges map[int64]known_uk_bridges.KnownBridge) {
	for _, underWay := range underways {
		if knownBridge, ok := knownBridges[underWay.Way.Id]; ok {
			underWay.IsExactKnownBridge = true
			underWay.KnownBridge = knownBridge
		}

		for _, neighbour := range underWay.Neighbours {
			if _, ok := knownBridges[neighbour.Id]; ok {
				underWay.HasNeighbouringKnownBridge = true
				underWay.KnownNeighbours = append(underWay.KnownNeighbours, neighbour)
				underWay.NeighboutMaxHeight = neighbour.MaxHeight
			}
		}
	}
}
