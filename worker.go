package datafeeder

import (
	//"fmt"
	"github.com/bndr/gojenkins"
  "time"
)

func IsWithinTimeframe(build *gojenkins.Build, timeframe int) bool {
	timeframeDuration := time.Duration(timeframe) * time.Second
	buildInfo := build.Info()
	jobEndTime := GetBuildEndTime(buildInfo.Timestamp / 1000, buildInfo.Duration / 1000, time.Second)
	diff := time.Now().Sub(jobEndTime)
	return !(diff >= timeframeDuration) && !buildInfo.Building
}


func ParseTimeStamp(timestamp int64) time.Time {
  return time.Unix(timestamp, 0)
}

// Adds the start time of the build with the duration to get the end time
func GetBuildEndTime(timestamp int64, duration int64, conv time.Duration) time.Time {
	return ParseTimeStamp(timestamp).Add(time.Duration(duration) * conv)
}
