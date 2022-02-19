package square

import "math"

type sidesNumber int

const (
	SidesCircle   = sidesNumber(0)
	SidesTriangle = sidesNumber(3)
	SidesSquare   = sidesNumber(4)
)

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
