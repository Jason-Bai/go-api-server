package user

import (
	. "go-api-server/demo09/handler"
	"go-api-server/demo09/model"
	"go-api-server/demo09/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
