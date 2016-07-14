package datafeeder

import (
	elastic "gopkg.in/olivere/elastic.v3"
  "log"
)

type Document struct {
  Client     *elastic.Client
	JsonBody   interface{}
	Index      string
	RecordType string
  Timestamp  string
	Id         string
	Refresh    bool
}

type Special struct {
  Name string `json:"name"`
}

func (d *Document) IndexDocument() error {
  log.Println(d.Timestamp)
  _, err := d.Client.Index().Index(d.Index).Type(d.RecordType).Timestamp(d.Timestamp).Id(d.Id).BodyJson(d.JsonBody).Refresh(d.Refresh).Do()
  log.Println("Indexed Document", d.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateElasticsearchIndex(client *elastic.Client, index string) error {
  mapping := `{
    "settings":{
        "number_of_shards":1,
        "number_of_replicas":0
    },
    "mappings":{
        "jenkinslogs":{
        "properties":{
          "timestamp": {"type": "date"}
        }
      }
    }
}`
	exists, err := client.IndexExists(index).Do()
	if err != nil {
		return err
	}
	if !exists {
		_, err = client.CreateIndex(index).BodyString(mapping).Do()
		if err != nil {
			return err
		}
	}
	return nil
}
