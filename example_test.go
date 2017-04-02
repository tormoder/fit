package fit_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tormoder/fit"
)

func Example() {
	// Read our FIT test file data
	testFile := filepath.Join("testdata", "fitsdk", "Activity.fit")
	testData, err := ioutil.ReadFile(testFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode the FIT file data
	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Inspect the TimeCreated field in the FileId message
	fmt.Println(fit.FileId.TimeCreated)

	// Inspect the dynamic Product field in the FileId message
	fmt.Println(fit.FileId.GetProduct())

	// Inspect the FIT file type
	fmt.Println(fit.Type())

	// Get the actual activity
	activity, err := fit.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the latitude and longitude of the first Record message
	for _, record := range activity.Records {
		fmt.Println(record.PositionLat)
		fmt.Println(record.PositionLong)
		break
	}

	// Print the sport of the first Session message
	for _, session := range activity.Sessions {
		fmt.Println(session.Sport)
		break
	}

	// Output:
	// 2012-04-09 21:22:26 +0000 UTC
	// Hrm1
	// Activity
	// 41.51393
	// -73.14859
	// Running
}
