package square

import (
	"math"
)

type intCustomType int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	const SidesCircle intCustomType = 0
	const SidesTriangle intCustomType = 3
	const SidesSquare intCustomType = 4
	sqrTriangle := (math.Sqrt(3) * sideLen * sideLen) / 4
	sqrSquare := sideLen * sideLen
	sqrCircle := math.Pi * sideLen * sideLen
	if sidesNum == SidesTriangle {
		return sqrTriangle
	} else if sidesNum == SidesSquare {
		return sqrSquare
	} else if sidesNum == SidesCircle {
		return sqrCircle
	} else {
		return 0
	}
}
