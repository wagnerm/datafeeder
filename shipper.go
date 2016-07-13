package datafeeder

import (
	"github.com/bndr/gojenkins"
  elastic "gopkg.in/olivere/elastic.v3"
)

func CreateShipperIndex(client *elastic.Client, index string) error {
  exists, err := client.IndexExists(index).Do()
  if err != nil {
    return err
  }
  if !exists {
    _, err = client.CreateIndex(index).Do()
  	if err != nil {
      return err
  	}
  }
  return nil
}

func ShipElasticsearch(client *elastic.Client, jobBuild *gojenkins.Build, index string, recordType string) error {
	_, err := client.Index().Index(index).Type(recordType).BodyJson(jobBuild.Info()).Refresh(true).Do()
	if err != nil {
    return err
	}
  return nil
}
