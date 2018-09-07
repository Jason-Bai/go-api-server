package user

import (
	. "go-api-server/demo14/handler"
	"go-api-server/demo14/model"
	"go-api-server/demo14/pkg/auth"
	"go-api-server/demo14/pkg/errno"
	"go-api-server/demo14/pkg/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
