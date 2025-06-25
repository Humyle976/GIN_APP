package helpers

import (
	"context"
	"errors"
	"gin_app/config"
)

func GetTigerGraphToken() (string, error) {
	ctx := context.Background()
	res, err := config.Client.Exists(ctx, "auth:tigergraph:token").Result()

	if err != nil {
		return "", errors.New("internal server error")
	}

	if res == 1 {
		res, err := config.Client.Get(ctx, "auth:tigergraph:token").Result()
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return config.ConnectTigerGraph()
}
