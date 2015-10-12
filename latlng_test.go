package fit

import (
	"math"
	"testing"
)

var latlngSemiTests = []struct {
	semi             int32
	latsemi, lngsemi int32
	latstr, lngstr   string
}{
	{
		0,
		0,
		0,
		"0.00000",
		"0.00000",
	},
	{
		703539217,
		703539217,
		703539217,
		"58.96997",
		"58.96997",
	},
	{
		68398630,
		68398630,
		68398630,
		"5.73311",
		"5.73311",
	},
	{
		-649405020,
		-649405020,
		-649405020,
		"-54.43250",
		"-54.43250",
	},

	{
		math.MaxInt32,
		sint32Invalid,
		sint32Invalid,
		"Invalid",
		"Invalid",
	},
	{
		math.MaxInt32 - 1,
		sint32Invalid,
		math.MaxInt32 - 1,
		"Invalid",
		"180.00000",
	},
	{
		math.MaxInt32 / 2,
		math.MaxInt32 / 2,
		math.MaxInt32 / 2,
		"90.00000",
		"90.00000",
	},
	{
		(math.MaxInt32 / 2) + 1,
		sint32Invalid,
		(math.MaxInt32 / 2) + 1,
		"Invalid",
		"90.00000",
	},
	{
		(math.MaxInt32 / 2) - 1,
		(math.MaxInt32 / 2) - 1,
		(math.MaxInt32 / 2) - 1,
		"90.00000",
		"90.00000",
	},

	{
		math.MinInt32,
		sint32Invalid,
		math.MinInt32,
		"Invalid",
		"-180.00000",
	},
	{
		math.MinInt32 / 2,
		math.MinInt32 / 2,
		math.MinInt32 / 2,
		"-90.00000",
		"-90.00000",
	},
	{
		(math.MinInt32 / 2) + 1,
		(math.MinInt32 / 2) + 1,
		(math.MinInt32 / 2) + 1,
		"-90.00000",
		"-90.00000",
	},
	{
		(math.MaxInt32 / 2) + 1,
		sint32Invalid,
		(math.MaxInt32 / 2) + 1,
		"Invalid",
		"90.00000",
	},
	{
		sint32Invalid,
		sint32Invalid,
		sint32Invalid,
		"Invalid",
		"Invalid",
	},
}

func TestLatLngFromSemicircles(t *testing.T) {
	for i, test := range latlngSemiTests {
		latFromSemi := NewLatitude(test.latsemi)
		if latFromSemi.Semicircles() != test.latsemi {
			t.Errorf("%d lat semi: got %d, want %d", i, latFromSemi.Semicircles(), test.latsemi)
		}
		if latFromSemi.String() != test.latstr {
			t.Errorf("%d lat str: got %q, want %q", i, latFromSemi.String(), test.latstr)
		}

		lngFromSemi := NewLongitude(test.semi)
		if lngFromSemi.Semicircles() != test.lngsemi {
			t.Errorf("%d lng semi: got %d, want %d", i, lngFromSemi.Semicircles(), test.lngsemi)
		}
		if lngFromSemi.String() != test.lngstr {
			t.Errorf("%d lng str: got %q, want %q", i, lngFromSemi.String(), test.lngstr)
		}
	}
}

var latlngDegTests = []struct {
	deg              float64
	latsemi, lngsemi int32
	latstr, lngstr   string
}{
	{
		0,
		0,
		0,
		"0.00000",
		"0.00000",
	},
	{
		58.96997,
		703539146,
		703539146,
		"58.96997",
		"58.96997",
	},
	{
		5.73311,
		68398666,
		68398666,
		"5.73311",
		"5.73311",
	},
	{
		-54.43250,
		-649405020,
		-649405020,
		"-54.43250",
		"-54.43250",
	},
	{
		180.00001,
		sint32Invalid,
		sint32Invalid,
		"Invalid",
		"Invalid",
	},
	{
		180.00000,
		sint32Invalid,
		sint32Invalid,
		"Invalid",
		"Invalid",
	},
	{
		179.99999,
		sint32Invalid,
		2147483528,
		"Invalid",
		"179.99998",
	},
	{
		90.00001,
		sint32Invalid,
		1073741943,
		"Invalid",
		"90.00001",
	},
	{
		90.00000,
		sint32Invalid,
		1073741824,
		"Invalid",
		"90.00000",
	},
	{
		89.99999,
		1073741704,
		1073741704,
		"89.99999",
		"89.99999",
	},
	{
		math.MaxFloat64,
		sint32Invalid,
		sint32Invalid,
		"Invalid",
		"Invalid",
	},
}

func TestLatLngFromDegrees(t *testing.T) {
	for i, test := range latlngDegTests {
		latFromDeg := NewLatitudeDegrees(test.deg)
		if latFromDeg.Semicircles() != test.latsemi {
			t.Errorf("%d lat semi: got %d, want %d", i, latFromDeg.Semicircles(), test.latsemi)
		}
		if latFromDeg.String() != test.latstr {
			t.Errorf("%d lat str: got %q, want %q", i, latFromDeg.String(), test.latstr)
		}

		lngFromDeg := NewLongitudeDegrees(test.deg)
		if lngFromDeg.Semicircles() != test.lngsemi {
			t.Errorf("%d lng semi: got %d, want %d", i, lngFromDeg.Semicircles(), test.lngsemi)
		}
		if lngFromDeg.String() != test.lngstr {
			t.Errorf("%d lng str: got %q, want %q", i, lngFromDeg.String(), test.lngstr)
		}
	}
}
