package router

import (
	"fmt"
	"math"
)

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

var DirectionText = map[DirectionType]string{
	Start:       "从起点出发",
	Straight:    "直行",
	SlightLeft:  "向左前方行驶",
	SlightRight: "向右前方行驶",
	Left:        "左转",
	Right:       "右转",
	SharpLeft:   "向左后方行驶",
	SharpRight:  "向右后方行驶",
}

type NavigationDirection struct {
	direction DirectionType
	way       string
	distance  float64
}

func (nd NavigationDirection) String() (s string) {
	if nd.way != "" {
		s = fmt.Sprintf("%s，进入<b>%s</b>，继续前行%.3f公里", DirectionText[nd.direction], nd.way, nd.distance)
	} else {
		s = fmt.Sprintf("%s，继续前行%.3f公里", DirectionText[nd.direction], nd.distance)
	}
	return
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
