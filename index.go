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

func (d *Document) IndexDocument() error {
	log.Println(d.Timestamp)
	_, err := d.Client.Index().Index(d.Index).Type(d.RecordType).Timestamp(d.Timestamp).Id(d.Id).BodyJson(d.JsonBody).Refresh(d.Refresh).Do()
	log.Println("Indexed Document", d.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateElasticsearchIndex(client *elastic.Client, index string, mapping string) error {
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
