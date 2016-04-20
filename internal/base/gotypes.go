package base

import "fmt"

// gotype is used in the profile field lookup table to represent the data type
// (or type category) for a field when decoded into a Go message struct.
type GoType uint8

const (
	Fit       GoType = 0 // Standard 	-> Fit base type/alias
	TimeUTC          = 1 // Time UTC 	-> time.Time
	TimeLocal        = 2 // Time Local 	-> time.Time with Location
	Lat              = 3 // Latitude 	-> fit.Latitude
	Lng              = 4 // Longitude 	-> fit.Longitude
)

func (g GoType) String() string {
	if int(g) >= len(gstring) {
		return fmt.Sprintf("<unknown base.GoType(%d)>", g)
	}
	return gstring[g]
}

var gstring = [...]string{
	"fit",
	"timeutc",
	"timelocal",
	"lat",
	"lng",
}

func (g GoType) InvalidValueAsString() string {
	if g == 0 {
		panic("no invalid string value for gotype of type fit")
	}
	if int(g) > len(tinvalid) {
		return fmt.Sprintf("<unknown base.GoType(%d)>", g)
	}
	return ginvalid[g]
}

var ginvalid = [...]string{
	"",
	"timeBase",
	"timeBase",
	"NewLatitudeInvalid()",
	"NewLongitudeInvalid()",
}
