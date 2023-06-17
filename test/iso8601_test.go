package test

import (
	"regexp"
	"testing"

	ISO8601 "github.com/averyyan/xz-time/iso8601"
)

func TestParseTimeDateISOZone(t *testing.T) {
	timedateString := "2023-10-10T00:00:00+08:00"
	timer, err := ISO8601.ParseTimeDateString(timedateString)
	t.Log(timer.Unix(), err)
}

type duration struct {
	Years   int
	Weeks   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func TestXXX(t *testing.T) {
	timedateString := "2023-10-10T00:00:00+08:00"
	full, err := regexp.Compile(`/^([\+-]?\d{4}(?!\d{2}\b))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([T\s]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$/`)
	t.Log(err)
	match := full.FindStringSubmatch(timedateString)
	t.Log(match)
}

func TestParseDuration(t *testing.T) {
	timedateString := "P15D"
	dur, err := ISO8601.ParseDuration(timedateString)
	t.Log(dur, err)
}
