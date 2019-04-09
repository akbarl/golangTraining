package square

import (
	"mypackages/package2"
)

func Area(a float64) float64 {
	return package2.Mul(a, a)
}

func Perimeter(a float64) float64 {
	return package2.Mul(a, 4)
}