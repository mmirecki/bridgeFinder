package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentIntersection(t *testing.T) {

	assert.True(t, Intersect(
		Segment{
			LatLng{1, 0},
			LatLng{0, 1},
		},
		Segment{
			LatLng{0, 0},
			LatLng{1, 1},
		},
	))

	assert.False(t, Intersect(
		Segment{
			LatLng{1, 0},
			LatLng{0, 1},
		},
		Segment{
			LatLng{0, 1},
			LatLng{1, 0},
		},
	))

	assert.False(t, Intersect(
		Segment{
			LatLng{1, 1},
			LatLng{0, 0},
		},
		Segment{
			LatLng{0, 0},
			LatLng{-1, -1},
		},
	))

	assert.False(t, Intersect(
		Segment{
			LatLng{1, 0},
			LatLng{0, 1},
		},
		Segment{
			LatLng{0.4, 0.4},
			LatLng{-1, -1},
		},
	))

	assert.False(t, Intersect(
		Segment{
			LatLng{0, 0},
			LatLng{1, 1},
		},
		Segment{
			LatLng{0, -1.1},
			LatLng{-1.1, 0},
		},
	))
}

func TestSegmentIntersection2(t *testing.T) {

	assert.False(t, Intersect(
		Segment{
			LatLng{-1.4604932, 50.9137268},
			LatLng{-1.4592071, 50.9149031},
		},
		Segment{
			LatLng{-1.4604932, 50.9137268},
			LatLng{-1.4609572, 50.9133024},
		},
	))

}

func TestSegmentIntersection3(t *testing.T) {

	assert.False(t, Intersect(
		Segment{
			LatLng{-1.3762777, 50.9177592},
			LatLng{-1.3768396, 50.9178226},
		},
		Segment{
			LatLng{-1.3762777, 50.9177592},
			LatLng{-1.3758352, 50.9177165},
		},
	))

	assert.False(t, Intersect(
		Segment{
			LatLng{-1.3762777, 50.9177592},
			LatLng{-1.3768396, 50.9178226},
		},
		Segment{
			LatLng{-1.3757312, 50.9177065},
			LatLng{-1.3758352, 50.9177165},
		},
	))

}

func TestSegmentIntersection4(t *testing.T) {

	assert.False(t, Intersect(
		Segment{
			LatLng{-1.3979497, 50.9068976},
			LatLng{-1.3979347, 50.9069949},
		},
		Segment{
			LatLng{-1.3979058, 50.9071336},
			LatLng{-1.3979347, 50.9069949},
		},
	))

}

 -1.4737231,50.9213949
 -1.4738529, 50.9214255


 -1.4733628, 50.9215595
-1.4768113, 50.9205907
func TestSegmentIntersection5(t *testing.T) {

	assert.False(t, Intersect(
		Segment{
			LatLng{ -1.4737231,50.9213949},
			LatLng{ -1.4738529, 50.9214255},
		},
		Segment{
			LatLng{ -1.4733628, 50.9215595},
			LatLng{-1.4768113, 50.9205907},
		},
	))

}
















func TestSegmentIntersectionX(t *testing.T) {

	assert.False(t, Intersect(
		Segment{
			LatLng{0, 0},
			LatLng{0, 0},
		},
		Segment{
			LatLng{0, 0},
			LatLng{0, 0},
		},
	))

}

