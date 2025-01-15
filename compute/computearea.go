package compute

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/mmirecki/bridgeFinder/known_uk_bridges"
	"github.com/mmirecki/bridgeFinder/reporting"
	"github.com/mmirecki/bridgeFinder/utils"
)

func ComputeArea(knownBridges map[int64]known_uk_bridges.KnownBridge, minLatLng, maxLatLnt data.LatLng) ([]*data.UnderWay, error) {

	combinedWays := []*data.UnderWay{}
	combinedStats := data.BatchStats{}
	// 51.4, -0.4, 51.6, 0.44
	lng := minLatLng.Lng
	for lngN := 0; lng < maxLatLnt.Lng; lngN++ {
		lng = minLatLng.Lng + float64(lngN)*utils.UK_LNG_INCREMENT
		endLng := lng + utils.UK_LNG_INCREMENT

		lat := minLatLng.Lat
		for latN := 0; lat < maxLatLnt.Lat; latN++ {

			lat = minLatLng.Lat + float64(latN)*utils.UK_LAT_INCREMENT
			endLat := lat + utils.UK_LAT_INCREMENT
			fmt.Printf("===================\n")
			fmt.Printf("Computing for: %d %d    LATLNG: %.2f, %.2f\n", lngN, latN, lat, lng)

			if !utils.IsLatLngInUk(data.LatLng{Lat: lat, Lng: lng}) {
				reporting.WriteNotInUkToFile(lngN, latN, lng, lat)
				fmt.Printf("  NOT IN UK\n")
				continue
			}

			ways, err := computeSquare(knownBridges, lat, lng, endLat, endLng)
			if err != nil {
				reporting.WriteErrorToFile(err, lngN, latN, lng, lat)
				continue
			}
			combinedWays = append(combinedWays, ways...)

			reporting.WriteReportToFiles(ways)
			stats := reporting.ComputeStats(ways)
			combinedStats = combinedStats.Add(stats)
			reporting.WriteDoneToFile(err, lngN, latN, lng, lat, stats)
		}
	}

	waysById := make(map[int64][]*data.UnderWay)

	for _, way := range combinedWays {
		if way.IsExactKnownBridge || !way.HasNeighbouringKnownBridge {
			continue
		}

		fmt.Printf("Way: %v\n", way)

		if _, ok := waysById[way.Way.Id]; !ok {
			waysById[way.Way.Id] = []*data.UnderWay{}
		}
		waysById[way.Way.Id] = append(waysById[way.Way.Id], way)
	}
	reporting.WriteByWayToFile(waysById)

	fmt.Printf("Combined stats: %+v\n", combinedStats)

	fmt.Printf("Count: %d\n", combinedStats.Count)
	fmt.Printf("Missing: %d\n", combinedStats.MissingCount)
	fmt.Printf("Known: %d\n", combinedStats.KnownCount)
	fmt.Printf("Has neighbour: %d\n", combinedStats.HasNeighbourCount)
	return combinedWays, nil
}
