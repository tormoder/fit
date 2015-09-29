package fit

import "time"

var timeBase = time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)

func decodeDateTime(dt uint32) time.Time {
	return timeBase.Add(time.Duration(dt) * time.Second)
}

func encodeTime(t time.Time) uint32 {
	return uint32(t.Sub(timeBase) / time.Second)
}
