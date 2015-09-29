package fit

import (
	"math"
	"testing"
)

var semicirclesTests = []struct {
	lat  Latitude
	lng  Longitude
	semi int32
}{
	{
		Latitude(0),
		Longitude(0),
		0,
	},
	{
		Latitude(58.969975942745805),
		Longitude(58.969975942745805),
		703539217,
	},
	{
		Latitude(5.733106937259436),
		Longitude(5.733106937259436),
		68398630,
	},
	{
		Latitude(61.6361899394542),
		Longitude(61.6361899394542),
		735348389,
	},
	{
		Latitude(8.312469916418195),
		Longitude(8.312469916418195),
		99171629,
	},
	{
		Latitude(-54.43249996751547),
		Longitude(-54.43249996751547),
		-649405020,
	},
	{
		Latitude(3.418399952352047),
		Longitude(3.418399952352047),
		40783100,
	},
	{
		Latitude(179.99999991618097),
		Longitude(179.99999991618097),
		math.MaxInt32,
	},
}

func TestSemicirclesEncDec(t *testing.T) {
	for i, sct := range semicirclesTests {
		semilat := sct.lat.toSemicircles()
		if semilat != sct.semi {
			t.Errorf("%d: got %d, want %d from Latitude toSemicircles", i, semilat, sct.semi)
		}
		lat := NewLatitudeSemicircles(sct.semi)
		if lat != sct.lat {
			t.Errorf("%d: , got %v, want %v from NewLatitudeSemicircles", i, lat, sct.semi)
		}
		semilng := sct.lng.toSemicircles()
		if semilng != sct.semi {
			t.Errorf("%d: got %d, want %d from Longitude toSemicircles", i, semilng, sct.semi)
		}
		lng := NewLongitudeSemicircles(sct.semi)
		if lat != sct.lat {
			t.Errorf("%d: , got %v, want %v from NewLongitudeSemicircles", i, lng, sct.semi)
		}
	}
}
