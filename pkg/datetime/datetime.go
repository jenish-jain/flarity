package datetime

import "time"

const DDMMYYYY = "02-01-2006"

func StringToDate(dateString string, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}
