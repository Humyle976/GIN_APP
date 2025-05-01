package helpers

import (
	"context"
	"errors"
	"gin_app/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if float64(time.Now().Unix()) > (exp) {
			return nil, errors.New("token has expired")
		}
	}

	ctx := context.Background()
	res, err := config.Client.SIsMember(ctx, "auth:tokens:blacklist", tokenStr).Result()

	if err != nil {
		return nil, errors.New("internal server error")
	}

	if res {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
