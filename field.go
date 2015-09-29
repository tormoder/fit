package fit

import "fmt"

// field represents a fit message field in the profile field lookup table.
type field struct {
	sindex int
	scale  float32
	offset int16
	array  uint8
	t      gotype
	num    byte
	btype  fitBaseType
}

func (f field) String() string {
	return fmt.Sprintf(
		"num: %d | btype: %v | sindex: %d | scale: %v | offset: %d | array: %d",
		f.num, f.btype, f.sindex, f.scale, f.offset, f.array,
	)
}

// gotype is used in the profile field lookup table to represent the data type
// (or type category) for a field when decoded into a Go message struct.
type gotype uint8

const (
	fit gotype = iota // Standard -> Fit base type/alias

	// Special (non-profile types)
	timeutc   // Time UTC 	-> time.Time
	timelocal // Time Local -> time.Time with Location
	lat       // Latitude 	-> fit.Latitude
	lng       // Longitude 	-> fit.Longitude
	float     // float64 	-> Implies scale (and offset)
)

func (g gotype) String() string {
	if int(g) > len(gotypeString) {
		return fmt.Sprintf("gotype(%d)", g)
	}
	return gotypeString[g]
}

var gotypeString = [...]string{
	"fit",
	"timeutc",
	"timelocal",
	"lat    ",
	"lng     ",
	"float",
}
