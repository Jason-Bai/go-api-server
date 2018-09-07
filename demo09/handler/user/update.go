package user

import (
	. "go-api-server/demo09/handler"
	"go-api-server/demo09/model"
	"go-api-server/demo09/pkg/errno"
	"go-api-server/demo09/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userID, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.ID = uint64(userID)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, u)
}
