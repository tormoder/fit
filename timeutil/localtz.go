package timeutil

import (
	"errors"
	"fmt"
	"time"

	"github.com/bradfitz/latlong"
	"github.com/tormoder/fit"
)

var (
	errNoActivityMsg      = errors.New("no activity messsage found")
	errNoGeoDataAvailable = errors.New("no geographical data available")
)

// SetLocalTimeZone use the external library github.com/bradfitz/latlong and
// available FIT file GPS data (Record, Session or Lap messages) to set the
// timezone of local timestamps. It processes FIT files of type "Activity" and
// "ActivitySummary". It currently only attempts to set the local time zone of
// the field "LocalTimestamp" for the file's "Activity" message.
// SetLocalTimeZone returns a nil error if a local time zone was set or a
// descriptive error otherwise.
func SetLocalTimeZone(fitFile *fit.Fit) error {
	if fitFile == nil {
		return errors.New("provided fit file was nil")
	}

	var (
		startLat    fit.Latitude
		startLng    fit.Longitude
		activityMsg *fit.ActivityMsg
	)

	switch fitFile.FileId.Type {

	case fit.FileActivity:
		a, err := fitFile.Activity()
		if err != nil {
			return err
		}

		if a.Activity == nil {
			return errNoActivityMsg
		}

		switch {
		case len(a.Records) > 0:
			startLat = a.Records[0].PositionLat
			startLng = a.Records[0].PositionLong
		case len(a.Sessions) > 0:
			startLat = a.Sessions[0].StartPositionLat
			startLng = a.Sessions[0].StartPositionLong
		case len(a.Laps) > 0:
			startLat = a.Laps[0].StartPositionLat
			startLng = a.Laps[0].StartPositionLong
		default:
			return errNoGeoDataAvailable
		}
		activityMsg = a.Activity

	case fit.FileActivitySummary:
		as, err := fitFile.ActivitySummary()
		if err != nil {
			return err
		}

		if as.Activity == nil {
			return errNoActivityMsg
		}

		switch {
		case len(as.Sessions) > 0:
			startLat = as.Sessions[0].StartPositionLat
			startLng = as.Sessions[0].StartPositionLong
		case len(as.Laps) > 0:
			startLat = as.Laps[0].StartPositionLat
			startLng = as.Laps[0].StartPositionLong
		default:
			return errNoGeoDataAvailable
		}
		activityMsg = as.Activity

	default:
		return fmt.Errorf(
			"can't set local time zone for provided fit file of type %v",
			fitFile.FileId.Type)

	}

	if fit.IsBaseTime(activityMsg.LocalTimestamp) {
		return errors.New("time stamp was set to fit base time (invalid)")
	}

	location, err := getLocalTimeZone(startLat, startLng)
	if err != nil {
		return nil
	}

	activityMsg.LocalTimestamp = activityMsg.LocalTimestamp.In(location)

	return nil
}

func getLocalTimeZone(lat fit.Latitude, lng fit.Longitude) (*time.Location, error) {
	if lat.Invalid() || lng.Invalid() {
		return nil, errNoGeoDataAvailable
	}

	ltz := latlong.LookupZoneName(lat.Degrees(), lng.Degrees())
	if ltz == "" {
		return nil, errors.New("found no timezone for provided geographical data")
	}

	location, err := time.LoadLocation(ltz)
	if err != nil {
		return nil, fmt.Errorf("error loading location for zone name %q: %v", ltz, err)
	}

	return location, nil
}
