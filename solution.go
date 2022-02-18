package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sides int

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
