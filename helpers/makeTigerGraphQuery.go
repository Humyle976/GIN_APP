package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func QueryTigerGraph(endpoint string, queryparams map[string]string) (interface{}, error) {
	host := os.Getenv("TG_HOST")
	graph := os.Getenv("TG_GRAPH")

	token, err := GetTigerGraphToken()

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/restpp/query/%s/%s", host, graph, endpoint)
	client := resty.New()
	res, err := client.R().SetHeader("Authorization", "Bearer "+token).SetQueryParams(queryparams).Get(url)

	if err != nil {
		return nil, err
	}

	var rawData map[string]interface{}
	json.Unmarshal([]byte(res.String()), &rawData)
	raw := rawData["results"].([]interface{})[0].(map[string]interface{})["RESULT"]

	return raw, nil
}
