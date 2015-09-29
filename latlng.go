package fit

import "math"

var (
	semiToDegFactor = 180 / math.Pow(2, 31)
	degToSemiFactor = math.Pow(2, 31) / 180
)

// Latitude represents the geographical coordinate latitude.
type Latitude float64

// NewLatitudeSemicircles returns a new latitude from a semicirle.
func NewLatitudeSemicircles(semicircles int32) Latitude {
	return Latitude(float64(semicircles) * semiToDegFactor)
}

func (l Latitude) toSemicircles() int32 {
	return int32(float64(l) * degToSemiFactor)
}

// Longitude represents the geographical coordinate longitude.
type Longitude float64

// NewLongitudeSemicircles returns a new longitude from a semicirle.
func NewLongitudeSemicircles(semicircles int32) Longitude {
	return Longitude(float64(semicircles) * semiToDegFactor)
}

func (l Longitude) toSemicircles() int32 {
	return int32(float64(l) * degToSemiFactor)
}
