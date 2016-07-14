package datafeeder

import (
	"time"
)

func GenUTCTimestampTag(timestamp int64) string {
	parsed := ParseTimeStamp(timestamp / 1000).UTC()
	return parsed.Format(time.RFC3339)
}
