package lib

import (
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateBearing(t *testing.T) {

	bearing := CalculateBearing(data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 0},
		End:   data.LatLng{Lat: 1, Lng: 0},
	})
	assert.Equal(t, 0.0, bearing)

	bearing = CalculateBearing(data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 0},
		End:   data.LatLng{Lat: 0, Lng: 1},
	})
	assert.Equal(t, 90.0, bearing)

	bearing = CalculateBearing(data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 0},
		End:   data.LatLng{Lat: -1, Lng: 0},
	})
	assert.Equal(t, 180.0, bearing)

	bearing = CalculateBearing(data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 0},
		End:   data.LatLng{Lat: 0, Lng: -1},
	})
	assert.Equal(t, 270.0, bearing)

	bearing = CalculateBearing(data.Segment{
		Start: data.LatLng{Lat: 34.0522, Lng: -118.2437},
		End:   data.LatLng{Lat: 40.7128, Lng: -74.0060},
	})
	assert.Equal(t, 65.91883966110919, bearing)

}
