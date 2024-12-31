package reporting

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/lib"
	"io"
	"os"
)

func WriteErrorToFile(reportedErr error, lngN int, latN int, lng, lat float64) error {
	file, err := os.OpenFile("ERROR_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f, %s\n", lngN, latN, lng, lat, reportedErr.Error()))
	return nil
}

func WriteNotInUkToFile(lngN int, latN int, lng, lat float64) error {
	file, err := os.OpenFile("NOT_IN_UK", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f\n", lngN, latN, lng, lat))
	return nil
}

func WriteDoneToFile(err error, lngN int, latN int, lng, lat float64, stats data.BatchStats) error {
	file, err := os.OpenFile("DONE_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d, %d, %f, %f,%d, %d,%d, %d\n", lngN, latN, lng, lat, stats.Count, stats.MissingCount, stats.KnownCount, stats.HasNeighbourCount))
	return nil
}

func WriteReportToFiles(completeUnderWays []*data.UnderWay) error {
	missingFile, err := os.OpenFile("MISSING_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	knownFile, err := os.OpenFile("KNOWN_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	knownNeighbourFile, err := os.OpenFile("KNOWN_NEIGHBOUR_FILE", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	resultsFile, err := os.OpenFile("MISSING_BRIDGES", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
					underWay.Overhead.Tags["name"],
				),
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
