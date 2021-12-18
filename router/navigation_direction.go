package router

import "math"

type DirectionType int

const (
	Start DirectionType = iota
	Straight
	SlightLeft
	SlightRight
	Left
	Right
	SharpLeft
	SharpRight
)

type NavigationDirection struct {
	direction DirectionType
	way       string
	distance  float64
}

func getDirection(prevBearing, curBearing float64) DirectionType {
	absDiff := math.Abs(curBearing - prevBearing)
	if absDiff <= 15.0 {
		return Straight
	}

	if (curBearing > prevBearing && absDiff < 180.0) || (curBearing < prevBearing && absDiff > 180.0) {
		if absDiff < 30.0 || absDiff > 330.0 {
			return SlightRight
		}
		if absDiff < 100.0 || absDiff > 260.0 {
			return Right
		}
		return SharpRight
	} else {
		if absDiff < 30.0 || absDiff > 330.0 {
			return SlightLeft
		}
		if absDiff < 100.0 || absDiff > 260.0 {
			return Left
		}
		return SharpLeft
	}
}
