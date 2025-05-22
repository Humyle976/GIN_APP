package helpers

import (
	"context"
	"errors"
	"gin_app/config"
	"log"
)

func GetTigerGraphToken() (string, error) {
	ctx := context.Background()
	res, err := config.Client.Exists(ctx, "auth:tigergraph:token").Result()

	if err != nil {
		log.Println("Couldnt get token")
		return "", errors.New("internal server error")
	}

	if res == 1 {
		res, err := config.Client.Get(ctx, "auth:tigergraph:token").Result()
		if err != nil {
			return "", err
		}
		log.Println("Returning token")
		return res, nil
	}

	return config.ConnectTigerGraph()
}
