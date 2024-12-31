package data

import "github.com/mmirecki/bridgeFinder/known_uk_bridges"

type LatLng struct {
	Lng float64
	Lat float64
}

type Node struct {
	Id  int64
	Lat float64
	Lng float64
}

type CameraPosition struct {
	Position  LatLng
	Heading   float64
	ImageLink string
}

type UnderWay struct {
	Way                 Way
	Overhead            Way
	IntersectingSegment Segment
	CameraPositions     []CameraPosition
	Results             []Result

	// Is this bridge in the set received from Routing?
	IsExactKnownBridge bool
	// The Way in the set received from Routing
	KnownBridge known_uk_bridges.KnownBridge

	// Is this bridge in the set received from Routing?
	HasNeighbouringKnownBridge bool
	// The Way in the set received from Routing
	//NeighbouringKnownBridge known_uk_bridges.KnownBridge

	KnownBridgeInProximity  bool
	KnownBridgesInProximity []known_uk_bridges.KnownBridge

	Neighbours      []*Way
	KnownNeighbours []*Way

	MaxHeight          string
	NeighboutMaxHeight string
}

type Result struct {
	ExistsResult string
	HeightResult string
	Position     CameraPosition
}

type Way struct {
	Id        int64
	Lat       float64
	Lng       float64
	NodesIds  []int64
	Nodes     []Node
	Tags      map[string]string
	MaxHeight string
}

type Element struct {
	ElementType string            `json:"type"`
	Id          int64             `json:"id"`
	Nodes       []int64           `json:"nodes"`
	Tags        map[string]string `json:"tags"`
	Lat         float64           `json:"lat"`
	Lng         float64           `json:"lon"`
}

type Segment struct {
	Start LatLng
	End   LatLng
}

func (s Segment) Reverse() Segment {
	return Segment{
		Start: s.End,
		End:   s.Start,
	}
}

func (w Way) IsBridge() bool {
	_, ok := w.Tags["bridge"]
	return ok
}

type BatchStats struct {
	Count             int
	MissingCount      int
	KnownCount        int
	HasNeighbourCount int
}

func (s BatchStats) Add(stats BatchStats) BatchStats {
	return BatchStats{
		Count:             s.Count + stats.Count,
		MissingCount:      s.MissingCount + stats.MissingCount,
		KnownCount:        s.KnownCount + stats.KnownCount,
		HasNeighbourCount: s.HasNeighbourCount + stats.HasNeighbourCount,
	}
}
