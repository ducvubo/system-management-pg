package initialize

import (
	"crypto/tls"
	"net/http"
	"sync"

	"system-management-pg/global"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

// ElasticsearchSingleton manages Elasticsearch connection
type ElasticsearchSingleton struct {
	client *elasticsearch.Client
}

var (
	esInstance *ElasticsearchSingleton
	esOnce     sync.Once
)

// GetElasticsearchInstance returns the singleton instance of Elasticsearch client
func GetElasticsearchInstance() *elasticsearch.Client {
	esOnce.Do(func() {
		esInstance = &ElasticsearchSingleton{}
		esInstance.initElasticsearch()
	})
	return esInstance.client
}

// initElasticsearch initializes the Elasticsearch connection
func (e *ElasticsearchSingleton) initElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{"https://160.187.229.179:9200"},
		Username:  "elastic",
		Password:  "Duc17052003*",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	checkErrorPanicElasticsearch(err, "Elasticsearch client initialization error")

	// Ping to check connection
	res, err := client.Ping()
	if err != nil {
		checkErrorPanicElasticsearch(err, "Elasticsearch ping error")
	}
	defer res.Body.Close()

	global.Logger.Info("Initialized Elasticsearch successfully with Singleton pattern")
	e.client = client
	global.EsClient = client
}

// CloseElasticsearch is a placeholder for cleanup
func (e *ElasticsearchSingleton) CloseElasticsearch() error {
	return nil
}

func checkErrorPanicElasticsearch(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// InitElasticsearch public function to initialize
func InitElasticsearch() {
	global.EsClient = GetElasticsearchInstance()
}