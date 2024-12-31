package lib

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
)

func FindWaysInArea(minLat, minLng, maxLat, maxLng float64) ([]data.Way, error) {

	query := fmt.Sprintf(`[out:json];
way(%f, %f, %f,  %f)[highway][highway!~"^(path|track|cycleway|footway|service|steps)$"][bridge];
(._;>;);
out;
`, minLat, minLng, maxLat, maxLng)

	//way, err := overpassWayQuery(query)
	elements, err := OverpassQuery[data.Element](query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	ways := ProcessElements(-1, elements)
	return ways, nil
}
