package lib

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
)

const StreetMapsLink = "http://maps.google.com/maps?q=&layer=c&cbll=%f,%f&cbp=11,%d,0,0,0"

func GetCameraPositionsForWay(crossRoad data.UnderWay) []data.CameraPosition {

	cameraPositions := []data.CameraPosition{}
	cameraAngleA := CalculateBearing(crossRoad.IntersectingSegment)
	cameraAngleB := CalculateBearing(crossRoad.IntersectingSegment.Reverse())

	positionA := data.CameraPosition{Position: crossRoad.IntersectingSegment.Start, Heading: cameraAngleA}
	positionB := data.CameraPosition{Position: crossRoad.IntersectingSegment.End, Heading: cameraAngleB}

	positionA.ImageLink = getLinkForPosition(positionA)
	positionB.ImageLink = getLinkForPosition(positionB)

	cameraPositions = append(cameraPositions, positionA)
	cameraPositions = append(cameraPositions, positionB)

	return cameraPositions
}

func getLinkForPosition(position data.CameraPosition) string {
	return fmt.Sprintf(StreetMapsLink, position.Position.Lat, position.Position.Lng, int(position.Heading))
}
