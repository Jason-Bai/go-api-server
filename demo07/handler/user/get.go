package user

import (
	. "rest/demo07/handler"
	"rest/demo07/model"
	"rest/demo07/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
