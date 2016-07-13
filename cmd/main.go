package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/wagnerm/datafeeder"
	//"log"
)

var (
	timeframe int = 3600
)

func main() {
	// create jenkins
	jenkins, err := gojenkins.CreateJenkins("http://localhost:8080/").Init()
	if err != nil {
		fmt.Println(err)
	}

	//get all jobs
	alljobs, err := jenkins.GetAllJobs()
	if err != nil {
		fmt.Println(err)
	}

	// Find recent builds for jobs
	for _, j := range alljobs {
		if enabled, _ := j.IsEnabled(); enabled == true {
			jobBuilds, err := collateBuilds(j)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(len(jobBuilds), jobBuilds)
			//jobBuilds = datafeeder.IsWithinTimeframe(jobBuilds, timeframe)
			//fmt.Println(len(jobBuilds), jobBuilds)
			/*
			builds := j.GetDetails().Builds
			for _, b := range builds {

        build, err := j.GetBuild(b.Number)
        if err != nil {
          fmt.Println(err)
        }
        if build == nil {
          log.Println("No build found", b.Number)
        }
				log.Println("Looking at build", b.Number)
				datafeeder.IsWithinTimeframe(build, timeframe)

			}*/
		}
	}
}

func collateBuilds(j *gojenkins.Job) ([]*gojenkins.Build, error) {
	builds := j.GetDetails().Builds
	//jobBuilds := make([]*gojenkins.Build, len(builds))
	var jobBuilds []*gojenkins.Build
	for _, b := range builds {
			build, err := j.GetBuild(b.Number)
			if err != nil {
				return nil, err
			}
			if datafeeder.IsWithinTimeframe(build, timeframe) {
				jobBuilds = append(jobBuilds, build)
			}
	}
	return jobBuilds, nil
}
