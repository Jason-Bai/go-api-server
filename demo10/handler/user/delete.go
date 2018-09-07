package user

import (
	. "go-api-server/demo10/handler"
	"go-api-server/demo10/model"
	"go-api-server/demo10/pkg/errno"
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
