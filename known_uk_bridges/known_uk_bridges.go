package known_uk_bridges

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed OUTUPUT_UK_bridges1
var KnownUkBridges string

type KnownBridge struct {
	Id     int64
	Lat    float64
	Lng    float64
	Height float64
}

func GetKnownUKBridges() map[int64]KnownBridge {

	scanner := bufio.NewScanner(strings.NewReader(KnownUkBridges))
	bridges := map[int64]KnownBridge{}
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		id, _ := strconv.ParseInt(tokens[0], 10, 64)
		height, _ := strconv.ParseFloat(tokens[1], 64)
		lat, _ := strconv.ParseFloat(tokens[2], 64)
		lng, _ := strconv.ParseFloat(tokens[3], 64)
		bridges[id] = KnownBridge{Id: id, Height: height, Lat: lat, Lng: lng}
	}
	return bridges
}
