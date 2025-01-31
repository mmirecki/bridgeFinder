package dataquery

import (
	_ "embed"
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/lib"
	"github.com/mmirecki/bridgeFinder/reporting"
)

//go:embed bridges_southampton_10m.json
var SouthamptonData10m string

//go:embed bridges_southampton_100m.json
var SouthamptonData100m string

//go:embed debug.json
var DebugData string

func SouthamptonFileQuery() ([]data.Element, error) {
	elements, err := lib.OverpassFile([]byte(SouthamptonData100m))
	return elements, err
}

func DebugFileQuery() ([]data.Element, error) {
	elements, err := lib.OverpassFile([]byte(DebugData))
	return elements, err
}

func CoordinatesUnderpassQuery(minLat, minLng, maxLat, maxLng float64) ([]data.Element, error) {
	query := `[out:json][timeout: 600];


(
  way(759226682);
)->.delimiter;

(
  way[bridge](%.2f, %.2f, %.2f, %.2f);
)->.bridges;

foreach .bridges -> .bridge {
  (.bridge;.bridge >;)->.bridge;
  (
  way(around.bridge:100)[highway][highway!~"^(path|track|cycleway|footway|service|steps|pedestrian|unclassified|proposed)$"][!bridge]; 
)->.underWays;
  (.underWays;.underWays >;)->.underWays;

  .underWays out body;
  .bridge out body;
  .delimiter out body;
}
`
	query = fmt.Sprintf(query, minLat, minLng, maxLat, maxLng)
	elements, body, err := lib.OverpassQuery(query)
	if err != nil {
		return nil, err
	}
	reporting.WriteOsmResults(body, minLng, minLat)
	return elements, err
}

type DataSet struct {
	elements       []data.Element
	elementPointer int64
}

type BridgeInputData struct {
	Bridge             data.Way
	PotentialUnderWays []data.Way
	UnderWays          []*data.UnderWay
}

func NewDataSetFromFiles(minLat, minLng, maxLat, maxLng float64) (*DataSet, error) {

	data, err := reporting.ReadOsmResults(minLng, minLat)
	if err != nil {
		return nil, err
	}
	elements, err := lib.ParseElements(data)
	if err != nil {
		return nil, err
	}
	return &DataSet{
		elements: elements,
	}, nil
}

func NewDataSetForBounds(minLat, minLng, maxLat, maxLng float64) (*DataSet, error) {

	elements, err := CoordinatesUnderpassQuery(minLat, minLng, maxLat, maxLng)

	//elements, err := SouthamptonFileQuery()
	if err != nil {
		return nil, err
	}
	return &DataSet{
		elements: elements,
	}, nil
}

func NewDataSet() (*DataSet, error) {

	elements, err := DebugFileQuery()
	if err != nil {
		return nil, err
	}
	return &DataSet{
		elements: elements,
	}, nil
}

func (d *DataSet) NextBridge() (*BridgeInputData, bool) {

	ways := []data.Way{}
	nodes := make(map[int64]data.Node)

	if d.elementPointer >= int64(len(d.elements)-1) {
		return &BridgeInputData{}, false
	}

	for {
		element := d.elements[d.elementPointer]
		d.elementPointer++

		if element.ElementType == lib.WAY {

			// This is a delimiter element that separates the bridge from the crossroads
			// This is hardcoded in the query
			if element.Id == 759226682 {
				return createBridgeInputData(ways, nodes), true
			}

			maxHeight := ""
			if height, ok := element.Tags["maxheight"]; ok {
				maxHeight = height
			}

			way := data.Way{
				Id:        element.Id,
				NodesIds:  element.Nodes,
				Tags:      element.Tags,
				MaxHeight: maxHeight,
			}

			ways = append(ways, way)

		} else if element.ElementType == lib.NODE {
			node := data.Node{
				Id:  element.Id,
				Lat: element.Lat,
				Lng: element.Lng,
			}
			nodes[element.Id] = node
		}
	}
}

func createBridgeInputData(ways []data.Way, nodes map[int64]data.Node) *BridgeInputData {

	if len(ways) == 0 {
		return nil
	}

	for i, way := range ways {
		for _, id := range way.NodesIds {
			way.Nodes = append(way.Nodes, nodes[id])
		}
		ways[i] = way
	}

	bridge := ways[len(ways)-1]
	crossRoads := ways[0 : len(ways)-1]

	return &BridgeInputData{
		Bridge:             bridge,
		PotentialUnderWays: crossRoads,
	}
}
