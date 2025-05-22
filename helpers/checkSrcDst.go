package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckSrcDst(c *gin.Context) {
	user, ok := GetUserFromContext(c)

	if !ok {
		return
	}

	userStr := strconv.FormatUint(uint64(user.ID), 10)
	dstStr := c.Param("id")
	if userStr == dstStr {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Can't send friend request to yourself",
		})
		return
	}
}
