package infrastructure

import (
	"log"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearch struct {
}

func NewElasticSearch() {
	es, err := elasticsearch8.NewClient(elasticsearch8.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
