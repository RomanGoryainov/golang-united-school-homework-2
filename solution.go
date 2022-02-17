package square

type intCustomType int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	const SidesCircle, SidesTriangle, SidesSquare intCustomType = 0, 3, 4
	if sidesNum == SidesCircle || sidesNum == SidesTriangle || sidesNum == SidesSquare {
		return 1
	} else {
		return 0
	}

}
