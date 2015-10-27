package timeutil_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tormoder/fit"
	"github.com/tormoder/fit/timeutil"
)

func ExampleSetLocalTimeZone() {
	testFile := filepath.Join("../testdata", "fitsdk", "Activity.fit")
	testData, err := ioutil.ReadFile(testFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		fmt.Println(err)
		return
	}

	activity, err := fit.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Post decode & pre SetLocalTimezone:")
	fmt.Println(activity.Activity.LocalTimestamp)
	err = timeutil.SetLocalTimeZone(fit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Post SetLocalTimezone:")
	fmt.Println(activity.Activity.LocalTimestamp)

	// Output:
	// Post decode & pre SetLocalTimezone:
	// 2012-04-09 17:24:51 -0400 FITLOCAL
	// Post SetLocalTimezone:
	// 2012-04-09 17:24:51 -0400 EDT
}
