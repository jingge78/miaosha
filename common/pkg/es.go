package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"miaosha-jjl/common/global"
	"miaosha-jjl/common/proto/product"
	"strconv"
	"sync"
)

func SetSearch(index string) error {
	search := `{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text",
                "analyzer": "ik_max_word"
            }
        }
    }
}`
	data, err := json.Marshal(search)
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  bytes.NewReader(data),
	}
	res, err := req.Do(context.Background(), global.Es)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return err

	}
	defer res.Body.Close()
	return nil
}
func CreateEs(products []*product.ProductDetailResponse) error {
	err := SetSearch("product")
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, productOne := range products {
		wg.Add(1)

		go func(productOne *product.ProductDetailResponse) {
			defer wg.Done()

			// Build the request body.
			data, err := json.Marshal(productOne)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "product",
				DocumentID: strconv.FormatInt(productOne.Id, 10),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), global.Es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), productOne.Id+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(productOne)
	}
	wg.Wait()
	return nil
}
