package lib

import (
	"errors"
	"github.com/mmirecki/bridgeFinder/data"
)

// Point represents a point in 2D space

// Segment represents a line segment defined by two points

// Orientation returns the orientation of the ordered triplet (p, q, r)
// 0 -> p, q and r are collinear
// 1 -> Clockwise
// 2 -> Counterclockwise
func Orientation(p, q, r data.LatLng) int {
	val := (q.Lng-p.Lng)*(r.Lat-p.Lat) - (q.Lat-p.Lat)*(r.Lng-p.Lng)
	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1
	}
	return 2
}

// OnSegment checks if point q lies on segment pr
func OnSegment(p, q, r data.LatLng) bool {
	return q.Lat <= max(p.Lat, r.Lat) && q.Lat >= min(p.Lat, r.Lat) &&
		q.Lng <= max(p.Lng, r.Lng) && q.Lng >= min(p.Lng, r.Lng)
}

// Intersect checks if two segments intersect
func Intersect(seg1, seg2 data.Segment) bool {

	p1, q1 := seg1.Start, seg1.End
	p2, q2 := seg2.Start, seg2.End

	if p1 == p2 || p1 == q2 || q1 == p2 || q1 == q2 {
		return false
	}

	o1 := Orientation(p1, q1, p2)
	o2 := Orientation(p1, q1, q2)
	o3 := Orientation(p2, q2, p1)
	o4 := Orientation(p2, q2, q1)

	// Special Cases
	// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	if o1 == 0 && OnSegment(p1, p2, q1) {
		return false
	}

	// p1, q1 and q2 are collinear and q2 lies on segment p1q1
	if o2 == 0 && OnSegment(p1, q2, q1) {
		return false
	}

	// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	if o3 == 0 && OnSegment(p2, p1, q2) {
		return false
	}

	// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	if o4 == 0 && OnSegment(p2, q1, q2) {
		return false
	}

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	return false
}

// Helper functions
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func FindIntersectionPoint(seg1, seg2 data.Segment) (data.LatLng, error) {
	dx1 := seg1.End.Lng - seg1.Start.Lng
	dy1 := seg1.End.Lat - seg1.Start.Lat
	dx2 := seg2.End.Lng - seg2.Start.Lng
	dy2 := seg2.End.Lat - seg2.Start.Lat

	denominator := dx1*dy2 - dy1*dx2

	if denominator == 0 {
		return data.LatLng{}, errors.New("segments are parallel or coincident")
	}

	numerator1 := (seg2.Start.Lng-seg1.Start.Lng)*dy2 - (seg2.Start.Lat-seg1.Start.Lat)*dx2
	numerator2 := (seg2.Start.Lng-seg1.Start.Lng)*dy1 - (seg2.Start.Lat-seg1.Start.Lat)*dx1

	t1 := numerator1 / denominator
	t2 := numerator2 / denominator

	if t1 < 0 || t1 > 1 || t2 < 0 || t2 > 1 {
		return data.LatLng{}, errors.New("segments do not intersect")
	}

	intersection := data.LatLng{
		Lat: seg1.Start.Lat + t1*dy1,
		Lng: seg1.Start.Lng + t1*dx1,
	}

	return intersection, nil
}
