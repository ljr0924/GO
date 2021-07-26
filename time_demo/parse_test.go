package time_demo

import (
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {

	parsedWithoutLocation, _ := time.Parse("20060102", "20210723")

	t.Logf("without localtion: %d", parsedWithoutLocation.Unix())

	parsedWithLocation, _ := time.ParseInLocation("20060102", "20210723", time.Local)

	t.Logf("without localtion: %d", parsedWithLocation.Unix())

}
