package utils

import "time"

func ToLocalTime(unix int64, offset int) time.Time { // unix: time, offset: timezone
	loc := time.FixedZone("local", offset)
	return time.Unix(unix, 0).In(loc)
}
