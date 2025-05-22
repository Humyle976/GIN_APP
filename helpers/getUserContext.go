package helpers

import (
	"encoding/json"
	"fmt"
	"gin_app/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func helpInitial(c *gin.Context) int {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		return
	}
	users := user.(models.User)
	userStr := strconv.FormatUint(uint64(users.ID), 10)

	dstStr := c.Param("id")

	if userStr == dstStr {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
	}

	token, err := helpers.GetTigerGraphToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to query",
		})
	}

	client := resty.New()

	host := os.Getenv("TG_HOST")
	graph := os.Getenv("TG_GRAPH")
	url := fmt.Sprintf("%s/restpp/query/%s/DeleteFriendRequest", host, graph)
	res, err := client.R().SetHeader("Authorization", "Bearer "+token).SetQueryParams(map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	}).Get(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Unable to make request",
		})
		return
	}
	var raw map[string]interface{}

	json.Unmarshal([]byte(res.String()), &raw)
	result := raw["results"].([]interface{})[0].(map[string]interface{})["STATUS_CODE"]

	status := result.(float64)
}
