package fit

import (
	"testing"
	"time"
)

var timeTests = []struct {
	secs uint32
	time time.Time
}{
	{0, time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)},
	{1, time.Date(1989, time.December, 31, 0, 0, 1, 0, time.UTC)},
	{60 * 60 * 24, time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)},
	{172800, time.Date(1990, time.January, 2, 0, 0, 0, 0, time.UTC)},
	{14515200, time.Date(1990, time.June, 17, 0, 0, 0, 0, time.UTC)},
	{771897600, time.Date(2014, time.June, 17, 0, 0, 0, 0, time.UTC)},
}

func TestDateTimeEncDec(t *testing.T) {
	for i, tt := range timeTests {
		decGot := decodeDateTime(tt.secs)
		if decGot != tt.time {
			t.Errorf("%d: want %v from DecodeDateTime, got %v", i, tt.time, decGot)
		}
		encGot := encodeTime(tt.time)
		if encGot != tt.secs {
			t.Errorf("%d: want %v from EncodeTime, got %v", i, tt.secs, encGot)
		}
	}
}
