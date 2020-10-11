package elasticsearch

import (
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

type esClient struct {
	client *elasticsearch7.Client
}

type esClientInterface interface {
	Init()
	GetClient() *elasticsearch7.Client
}

//ElasticClient is the client for elastic db
var (
	ElasticClient esClientInterface = &esClient{}
)

func (c *esClient) Init() {
	var err error
	c.client, err = elasticsearch7.NewDefaultClient()
	if err != nil {
		panic(err)
	}
}

func (c *esClient) GetClient() *elasticsearch7.Client {
	return c.client
}
