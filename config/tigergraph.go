package config

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TokenResponse struct {
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
	Error      bool   `json:"error"`
	Message    string `json:"message"`
}

func ConnectTigerGraph() (string, error) {

	host := os.Getenv("TG_HOST")
	graph := os.Getenv("TG_GRAPH")
	username := os.Getenv("TG_USER")
	password := os.Getenv("TG_PASS")
	url := fmt.Sprintf("%s/gsql/v1/tokens", host)

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

	body := map[string]string{"graph": graph}
	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("TigerGraph token request failed: %s", resp.Status)
	}

	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}

	expiry, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", tr.Expiration)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	_, err = Client.SetNX(ctx, "auth:tigergraph:token", tr.Token, time.Until(expiry)).Result()
	if err != nil {
		return "", err
	}

	log.Println("TigerGraph token cached.")
	return tr.Token, nil
}
