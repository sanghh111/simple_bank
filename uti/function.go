package uti

import "time"

func StringToDateTime(dateTimeString string) (time.Time, error) {
	return time.Parse(DateTimeLayout, dateTimeString)
}

func DateTimeToString(dateTime time.Time) string {
	return dateTime.String()
}
