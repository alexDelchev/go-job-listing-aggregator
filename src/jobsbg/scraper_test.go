// +build test

package jobsbg

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatPublishingDateString(t *testing.T) {
	const dateLayout string = "02.01.06"

	expectedToday := time.Now().Format(dateLayout)
	expectedYesterday := time.Now().AddDate(0, 0, -1).Format(dateLayout)
	expectedLastWeek := time.Now().AddDate(0, 0, -7).Format(dateLayout)

	// днес == today in bulgarian
	actualToday := formatPublishingDateString("днес")

	// вчера == yesterday in bulgarian
	actualYesterday := formatPublishingDateString("вчера")

	actualLastWeek := formatPublishingDateString(expectedLastWeek)

	var errorMesssagePrefix string = fmt.Sprintf("%s failed:", t.Name())
	if expectedToday != actualToday {
		t.Errorf("%s Expected today %s, got %s",
			errorMesssagePrefix, expectedToday, actualToday)
	}

	if expectedYesterday != actualYesterday {
		t.Errorf("%s Expected yesterday %s, got %s",
			errorMesssagePrefix, expectedYesterday, actualYesterday)
	}

	if expectedLastWeek != actualLastWeek {
		t.Errorf("%s Expected last week %s, got %s",
			errorMesssagePrefix, expectedLastWeek, actualLastWeek)
	}
}
