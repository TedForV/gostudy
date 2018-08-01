package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

var indexName = "test"

func main() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	existService := client.IndexExists(indexName)
	ok, err := existService.Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !ok {
		createService := client.CreateIndex(indexName)
		result, err := createService.Do(context.Background())
		if err != nil {
			panic(err)
		}
		if !result.Acknowledged {
			fmt.Println(result)
			return
		}
	}

}
