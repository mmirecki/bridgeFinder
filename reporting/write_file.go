package reporting

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/lib"
	"io"
	"os"
	"strings"
)

func WriteOsmResults(data []byte, lng, lat float64) error {
	fileName := fmt.Sprintf("%.2f_%.2f.osm_data", lat, lng)
	file, err := os.OpenFile("OUTPUT/OSM_DATA/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(string(data))
	return nil
}

func WriteErrorToFile(reportedErr error, lngN int, latN int, lng, lat float64) error {
	file, err := os.OpenFile("OUTPUT/results/ERROR_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f, %s\n", lngN, latN, lng, lat, reportedErr.Error()))
	return nil
}

func WriteNotInUkToFile(lngN int, latN int, lng, lat float64) error {
	file, err := os.OpenFile("OUTPUT/results/NOT_IN_UK", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f\n", lngN, latN, lng, lat))
	return nil
}

func WriteDoneToFile(err error, lngN int, latN int, lng, lat float64, stats data.BatchStats) error {
	file, err := os.OpenFile("OUTPUT/results/DONE_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f,%d, %d,%d, %d\n", lngN, latN, lng, lat, stats.Count, stats.MissingCount, stats.KnownCount, stats.HasNeighbourCount))
	return nil
}

func WriteReportToFiles(completeUnderWays []*data.UnderWay) error {
	missingFile, err := os.OpenFile("OUTPUT/results/MISSING_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	knownFile, err := os.OpenFile("OUTPUT/results/KNOWN_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	knownNeighbourFile, err := os.OpenFile("OUTPUT/results/KNOWN_NEIGHBOUR_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	resultsFile, err := os.OpenFile("OUTPUT/results/MISSING_BRIDGES", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer missingFile.Close()
	defer knownFile.Close()
	defer knownNeighbourFile.Close()
	defer resultsFile.Close()

	missingCount := 0
	knownCount := 0
	hasNeighbourCount := 0

	for _, cr := range completeUnderWays {
		cr.CameraPositions = lib.GetCameraPositionsForWay(*cr)
	}

	//"===== MISSING BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && !underWay.HasNeighbouringKnownBridge {
			missingCount++
			FileDump(missingFile, underWay, false)

			resultsFile.WriteString(
				fmt.Sprintf("%d, %f, %f, %s, %s, %d, %s \n",
					underWay.Way.Id,
					underWay.IntersectingSegment.Start.Lng,
					underWay.IntersectingSegment.Start.Lat,
					underWay.Way.Tags["name"],
					underWay.MaxHeight,
					underWay.Overhead.Id,
					underWay.Overhead.Tags["name"]),
			)
		}
	}

	//"\n\n\n===== KNOWN NEIGHBOUR BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && underWay.HasNeighbouringKnownBridge {
			hasNeighbourCount++
			FileDump(knownNeighbourFile, underWay, true)
		}
	}

	//==== KNOWN BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if underWay.IsExactKnownBridge {
			knownCount++
			FileDump(knownFile, underWay, false)
		}
	}

	fmt.Printf("Underways len: %v\n", len(completeUnderWays))
	fmt.Printf("Known count: %v\n", knownCount)
	fmt.Printf("Neighbor known count: %v\n", hasNeighbourCount)
	fmt.Printf("Missing count: %v\n", missingCount)
	return nil

}

func WriteByWayToFile(waysById map[int64][]*data.UnderWay) error {
	file, err := os.OpenFile("OUTPUT/results/OUTPUT_BY_WAYID", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	verboseFile, err := os.OpenFile("OUTPUT/results/OUTPUT_BY_WAYID_VERBOSE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer file.Close()
	defer verboseFile.Close()

	file.WriteString("WayId;<Bridge Id, Intersection Lat, Intersection Lng, Height, NodesIds ...; > \n")

	for wayId, ways := range waysById {
		//file.WriteString(fmt.Sprintf("WayId: %d\n", wayId))
		FileDumpByWayId(file, wayId, ways)
		FileDumpByWayId_Verbose(verboseFile, ways)
	}
	return nil

}
func FileDumpByWayId(writer io.StringWriter, wayId int64, ways []*data.UnderWay) {
	wayStrings := []string{}

	for _, way := range ways {

		nodes := []string{}
		for _, node := range way.Way.Nodes {
			nodes = append(nodes, fmt.Sprintf("%d", node.Id))
		}
		nodesString := strings.Join(nodes, ",")

		wayStrings = append(wayStrings, fmt.Sprintf(" %d, %f, %f, %s, %s", way.Overhead.Id, way.IntersectionPoint.Lat, way.IntersectionPoint.Lng, way.MaxHeight, nodesString))

	}
	wayString := fmt.Sprintf("%d; %s", wayId, strings.Join(wayStrings, ";"))
	writer.WriteString(wayString)
}

func FileDumpByWayId_Verbose(writer io.StringWriter, underWaysById []*data.UnderWay) {
	writer.WriteString("-------------------------------------------\n")
	writer.WriteString(fmt.Sprintf("%d \n", underWaysById[0].Way.Id))
	writer.WriteString(fmt.Sprintf("https://www.openstreetmap.org/way/%d \n", underWaysById[0].Way.Id))

	for _, underWay := range underWaysById {
		writer.WriteString(fmt.Sprintf("  BRIDGE: %d %s HEIGHT: %s  \n", underWay.Overhead.Id, underWay.Overhead.Tags["name"], underWay.MaxHeight))
		writer.WriteString(fmt.Sprintf("      https://www.openstreetmap.org/way/%d  \n", underWay.Overhead.Id))
		writer.WriteString(fmt.Sprintf("      StreetView:  \n"))
		for _, position := range underWay.CameraPositions {
			writer.WriteString(fmt.Sprintf("        %v  \n", position.Position))
			writer.WriteString(fmt.Sprintf("          %s  \n", position.ImageLink))
		}
	}

}

func FileDump(writer io.StringWriter, underWay *data.UnderWay, printNeighbours bool) {

	writer.WriteString("-------------------------------------------\n")
	writer.WriteString(fmt.Sprintf("  BRIDGE: %d %s   %s\n", underWay.Overhead.Id, underWay.Overhead.Tags["name"], fmt.Sprintf("https://www.openstreetmap.org/way/%d", underWay.Overhead.Id)))
	writer.WriteString(fmt.Sprintf("     Underway %d \"%s\" Height:\"%s\"  %s\n", underWay.Way.Id, underWay.Way.Tags["name"], underWay.MaxHeight, fmt.Sprintf("https://www.openstreetmap.org/way/%d", underWay.Way.Id)))
	if printNeighbours {
		for _, neighbour := range underWay.KnownNeighbours {
			writer.WriteString(fmt.Sprintf("          KNOWN Neighbour %d \"%s\" Height:\"%s\"  %s\n", neighbour.Id, neighbour.Tags["name"], neighbour.MaxHeight, fmt.Sprintf("https://www.openstreetmap.org/way/%d", neighbour.Id)))
		}
	}

	writer.WriteString(fmt.Sprintf("           StreetView: "))
	for _, position := range underWay.CameraPositions {
		writer.WriteString(fmt.Sprintf("              %+v %s  \n", position.Position, position.ImageLink))
	}
	writer.WriteString("\n")
}

func ComputeStats(completeUnderWays []*data.UnderWay) data.BatchStats {

	count := len(completeUnderWays)
	missingCount := 0
	knownCount := 0
	hasNeighbourCount := 0

	// Miising
	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && !underWay.HasNeighbouringKnownBridge {
			missingCount++
		}
	}

	//==== KNOWN NEIGHBOUR BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if !underWay.IsExactKnownBridge && underWay.HasNeighbouringKnownBridge {
			hasNeighbourCount++
		}
	}

	//===== KNOWN BRIDGES ===========\n")
	for _, underWay := range completeUnderWays {
		if underWay.IsExactKnownBridge {
			knownCount++
		}
	}

	return data.BatchStats{
		Count:             count,
		MissingCount:      missingCount,
		KnownCount:        knownCount,
		HasNeighbourCount: hasNeighbourCount,
	}

}
