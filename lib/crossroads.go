package lib

import (
	"github.com/mmirecki/bridgeFinder/data"
)

const NODE = "node"
const WAY = "way"

func ExtractCrossRoads(bridgeWay data.Way, ways []data.Way) ([]*data.UnderWay, error) {
	result := []*data.UnderWay{}
outer:
	for _, w := range ways {
		//nodeA := w.Nodes[len(w.Nodes)-1]

		for i := 0; i < len(w.Nodes)-1; i++ {
			nodeA := w.Nodes[i]
			nodeB := w.Nodes[i+1]

			segment := data.Segment{data.LatLng{Lat: nodeA.Lat, Lng: nodeA.Lng}, data.LatLng{Lat: nodeB.Lat, Lng: nodeB.Lng}}

			for i := 0; i < len(bridgeWay.Nodes)-1; i++ {
				bridgeNodeA := bridgeWay.Nodes[i]
				bridgeNodeB := bridgeWay.Nodes[i+1]

				bridgeSegement := data.Segment{data.LatLng{Lat: bridgeNodeA.Lat, Lng: bridgeNodeA.Lng}, data.LatLng{Lat: bridgeNodeB.Lat, Lng: bridgeNodeB.Lng}}

				if Intersect(bridgeSegement, segment) {
					intersectionPoint, err := FindIntersectionPoint(bridgeSegement, segment)
					if err != nil {
						return nil, err
					}
					result = append(result, &data.UnderWay{Way: w, Overhead: bridgeWay, IntersectingSegment: segment, IntersectingBridgeSegment: bridgeSegement, MaxHeight: w.MaxHeight, IntersectionPoint: intersectionPoint})

					continue outer
				}
			}
		}
	}
	return result, nil
}

func getBoundingBox(way data.Way) (float64, float64, float64, float64) {

	minLat, maxLat, minLng, maxLng := +90.0, -90.0, +180.0, -180.0

	for _, n := range way.Nodes {

		if n.Lat > maxLat {
			maxLat = n.Lat
		}
		if n.Lat < minLat {
			minLat = n.Lat
		}
		if n.Lng > maxLng {
			maxLng = n.Lng
		}
		if n.Lng < minLng {
			minLng = n.Lng
		}

	}
	return minLat, maxLat, minLng, maxLng

}

func ProcessElements(queriedWayId int64, elements []data.Element) []data.Way {

	ways := []data.Way{}
	nodes := make(map[int64]data.Node)

	for _, e := range elements {
		if e.ElementType == WAY {
			if e.Id == queriedWayId {
				continue
			}
			way := data.Way{
				Id:       e.Id,
				Lat:      e.Lat,
				Lng:      e.Lng,
				NodesIds: e.Nodes,
				Tags:     e.Tags,
			}
			ways = append(ways, way)
		} else if e.ElementType == NODE {
			node := data.Node{
				Id:  e.Id,
				Lat: e.Lat,
				Lng: e.Lng,
			}
			nodes[e.Id] = node
		}
	}

	for i, way := range ways {
		for _, id := range way.NodesIds {
			way.Nodes = append(way.Nodes, nodes[id])
		}
		ways[i] = way
	}
	return ways
}
