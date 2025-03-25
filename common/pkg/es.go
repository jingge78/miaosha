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
	for _, productO := range products {
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
		}(productO)
	}
	wg.Wait()
	return nil
}

func EsSearchByKeyWord(keyWord string) ([]*product.ProductDetailResponse, error) {
	var buf bytes.Buffer
	var r map[string]interface{}
	var query map[string]interface{}
	if keyWord != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"StoreName": keyWord,
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	// Perform the search request.
	res, err := global.Es.Search(
		global.Es.Search.WithContext(context.Background()),
		global.Es.Search.WithIndex("product"),
		global.Es.Search.WithBody(&buf),
		global.Es.Search.WithTrackTotalHits(true),
		global.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	var products []*product.ProductDetailResponse
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		products = append(products, &product.ProductDetailResponse{
			Id:        int64(hit.(map[string]interface{})["_source"].(map[string]interface{})["Id"].(float64)),
			Image:     hit.(map[string]interface{})["_source"].(map[string]interface{})["Image"].(string),
			StoreName: hit.(map[string]interface{})["_source"].(map[string]interface{})["StoreName"].(string),
			StoreInfo: hit.(map[string]interface{})["_source"].(map[string]interface{})["StoreInfo"].(string),
			Price:     float32(hit.(map[string]interface{})["_source"].(map[string]interface{})["Price"].(float64)),
			Postage:   float32(hit.(map[string]interface{})["_source"].(map[string]interface{})["Stock"].(float64)),
			IsPostage: int64(hit.(map[string]interface{})["_source"].(map[string]interface{})["IsPostage"].(float64)),
			Stock:     int64(hit.(map[string]interface{})["_source"].(map[string]interface{})["Stock"].(float64)),
			Browse:    int64(hit.(map[string]interface{})["_source"].(map[string]interface{})["Browse"].(float64)),
		})
	}
	return products, nil
}
