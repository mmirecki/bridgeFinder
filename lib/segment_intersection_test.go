package lib

import (
	"fmt"
	"github.com/mmirecki/bridgeFinder/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentIntersection(t *testing.T) {

	assert.True(t, Intersect(
		data.Segment{
			data.LatLng{1, 0},
			data.LatLng{0, 1},
		},
		data.Segment{
			data.LatLng{0, 0},
			data.LatLng{1, 1},
		},
	))

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{1, 0},
			data.LatLng{0, 1},
		},
		data.Segment{
			data.LatLng{0, 1},
			data.LatLng{1, 0},
		},
	))

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{1, 1},
			data.LatLng{0, 0},
		},
		data.Segment{
			data.LatLng{0, 0},
			data.LatLng{-1, -1},
		},
	))

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{1, 0},
			data.LatLng{0, 1},
		},
		data.Segment{
			data.LatLng{0.4, 0.4},
			data.LatLng{-1, -1},
		},
	))

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{0, 0},
			data.LatLng{1, 1},
		},
		data.Segment{
			data.LatLng{0, -1.1},
			data.LatLng{-1.1, 0},
		},
	))
}

func TestSegmentIntersection2(t *testing.T) {

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{-1.4604932, 50.9137268},
			data.LatLng{-1.4592071, 50.9149031},
		},
		data.Segment{
			data.LatLng{-1.4604932, 50.9137268},
			data.LatLng{-1.4609572, 50.9133024},
		},
	))

}

func TestSegmentIntersection3(t *testing.T) {

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{-1.3762777, 50.9177592},
			data.LatLng{-1.3768396, 50.9178226},
		},
		data.Segment{
			data.LatLng{-1.3762777, 50.9177592},
			data.LatLng{-1.3758352, 50.9177165},
		},
	))

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{-1.3762777, 50.9177592},
			data.LatLng{-1.3768396, 50.9178226},
		},
		data.Segment{
			data.LatLng{-1.3757312, 50.9177065},
			data.LatLng{-1.3758352, 50.9177165},
		},
	))

}

func TestSegmentIntersection4(t *testing.T) {

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{-1.3979497, 50.9068976},
			data.LatLng{-1.3979347, 50.9069949},
		},
		data.Segment{
			data.LatLng{-1.3979058, 50.9071336},
			data.LatLng{-1.3979347, 50.9069949},
		},
	))

}

func TestSegmentIntersection5(t *testing.T) {

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{-1.4737231, 50.9213949},
			data.LatLng{-1.4738529, 50.9214255},
		},
		data.Segment{
			data.LatLng{-1.4733628, 50.9215595},
			data.LatLng{-1.4768113, 50.9205907},
		},
	))

}

func TestSegmentIntersectionX(t *testing.T) {

	assert.False(t, Intersect(
		data.Segment{
			data.LatLng{0, 0},
			data.LatLng{0, 0},
		},
		data.Segment{
			data.LatLng{0, 0},
			data.LatLng{0, 0},
		},
	))
}

func TestFindIntersectionPoint(t *testing.T) {
	seg1 := data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 0},
		End:   data.LatLng{Lat: 1, Lng: 1},
	}
	seg2 := data.Segment{
		Start: data.LatLng{Lat: 0, Lng: 1},
		End:   data.LatLng{Lat: 1, Lng: 0},
	}
	point, err := FindIntersectionPoint(seg1, seg2)
	assert.Nil(t, err)

	assert.Equal(t, data.LatLng{0.5, 0.5}, point)
}
func TestFindIntersectionPoint2(t *testing.T) {
	seg1 := data.Segment{
		Start: data.LatLng{Lat: 37.7749, Lng: -122.4194},
		End:   data.LatLng{Lat: 34.0522, Lng: -118.2437},
	}
	seg2 := data.Segment{
		Start: data.LatLng{Lat: 36.7783, Lng: -119.4179},
		End:   data.LatLng{Lat: 32.7157, Lng: -117.1611},
	}

	intersection, err := FindIntersectionPoint(seg1, seg2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Intersection point: Lat %.6f, Lng %.6f\n", intersection.Lat, intersection.Lng)
	}
}
