package datafeeder

import (
	"github.com/bndr/gojenkins"
	"time"
)

// Check if a build has finished within timeframe and is not currently building
func IsWithinTimeframe(build *gojenkins.Build, timeframe int) bool {
	timeframeDuration := time.Duration(timeframe) * time.Second
	buildInfo := build.Info()
	jobEndTime := GetBuildEndTime(buildInfo.Timestamp/1000, buildInfo.Duration/1000, time.Second)
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

func CollateBuilds(j *gojenkins.Job, timeframe int) ([]*gojenkins.Build, error) {
	builds := j.GetDetails().Builds
	//jobBuilds := make([]*gojenkins.Build, len(builds)) // interesting that is produces nil values in the array
	// leaving here so I remember
	var jobBuilds []*gojenkins.Build
	for _, b := range builds {
		build, err := j.GetBuild(b.Number)
		if err != nil {
			return nil, err
		}
		if IsWithinTimeframe(build, timeframe) {
			jobBuilds = append(jobBuilds, build)
		}
	}
	return jobBuilds, nil
}

func GenUTCTimestampTag(timestamp int64) string {
	parsed := ParseTimeStamp(timestamp / 1000).UTC()
	return parsed.Format(time.RFC3339)
}
