package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/wagnerm/datafeeder"
	elastic "gopkg.in/olivere/elastic.v3"
	"strconv"
)

var (
	index                  = "jenkinslogs"
	timeframe              = 3600
	jenkinsDocumentMapping = `{
    	"mappings":{
        	"%s":{
        			"properties":{
          				"timestamp": {"type": "date"}
        			}
      		}
    	}
		}`
)

func main() {

	// create jenkins
	jenkins, err := gojenkins.CreateJenkins("http://localhost:8080/").Init()
	if err != nil {
		fmt.Println(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println(err)
	}
	err = datafeeder.CreateElasticsearchIndex(client, index, fmt.Sprint(jenkinsDocumentMapping, index))
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
		jobBuilds := make([]*gojenkins.Build, 0)
		if enabled, _ := j.IsEnabled(); enabled == true {
			jobBuilds, err = datafeeder.CollateBuilds(j, timeframe)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(len(jobBuilds), jobBuilds)
		for _, build := range jobBuilds {
			d := datafeeder.Document{Client: client, JsonBody: build.Info(), Index: index, RecordType: index, Timestamp: datafeeder.GenUTCTimestampTag(build.Info().Timestamp), Id: strconv.FormatInt(build.Info().Number, 10), Refresh: false}
			err = d.IndexDocument()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
