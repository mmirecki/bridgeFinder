package utils

import (
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLatLngInUk(t *testing.T) {

	result := IsLatLngInUk(data.LatLng{Lat: 55.20, Lng: 1.70})
	assert.True(t, result)

}
