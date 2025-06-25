package controllers

import (
	"gin_app/dto"
	"gin_app/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFriendsOfCurrentUser(c *gin.Context) {
	user, ok := helpers.GetUserFromContext(c)
	if !ok {
		return
	}

	userStr := strconv.FormatUint(uint64(user.ID), 10)
	res, err := helpers.QueryTigerGraph("/GetFriends", map[string]string{
		"userID": userStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	result := res.([]interface{})
	var friendlist []any
	for _, v := range result {
		attrs := v.(map[string]interface{})["attributes"].(map[string]interface{})

		age := attrs["age"].(float64)
		name := attrs["name"].(string)
		userID := attrs["user_id"].(float64)

		friendlist = append(friendlist, dto.GetFriendListResponseDTO(age, name, userID))
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": friendlist,
	})
}

func SendFriendRequest(c *gin.Context) {

	userStr, dstStr, ok := helpers.CheckSrcDst(c)

	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("SendFriendRequest", map[string]string{
		"src": userStr,
		"dst": dstStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 400 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User not found",
		})
		return
	} else if status == 40901 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Friend request already sent",
		})
		return
	} else if status == 40902 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Friend request already received",
		})
		return
	} else if status == 40903 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Already friends with the user",
		})
		return
	} else if status == 40301 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User is blocked",
		})
		return
	} else if status == 40302 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "You are blocked by the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Friend request sent",
	})
}

func AcceptFriendRequest(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)

	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("AcceptFriendRequest", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}
	status := result.(float64)

	if status == 400 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not found",
		})
		return
	} else if status == 409 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Already friends with the user",
		})
		return
	} else if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Friend request not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Friend request accepted",
	})
}

func DeclineFriendRequest(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)

	if !ok {
		return
	}
	result, err := helpers.QueryTigerGraph("DeclineFriendRequest", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Request not found",
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteFriendRequest(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)
	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("DeleteFriendRequest", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Friend request not found",
		})
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func RemoveAFriend(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)
	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("BlockAUser", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User is not a friend",
		})
		return
	}

	if status == 204 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

}
func BlockAUser(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)
	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("BlockAUser", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 409 {
		c.JSON(http.StatusConflict,gin.H{})
		return
	}

	if status == 204 {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func UnblockAUser(c *gin.Context) {
	userStr, dstStr, ok := helpers.CheckSrcDst(c)
	if !ok {
		return
	}

	result, err := helpers.QueryTigerGraph("BlockAUser", map[string]string{
		"srcID": userStr,
		"dstID": dstStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't send request",
		})
		return
	}

	status := result.(float64)

	if status == 409 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User already blocked",
		})
		return
	}

	if status == 204 {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
