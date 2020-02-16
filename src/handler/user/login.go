package user

import (
	"fmt"
	. "api-server-demo/src/handler"
	"api-server-demo/src/model"
	"api-server-demo/src/auth"
	"api-server-demo/src/errno"
	"github.com/gin-gonic/gin"
)

// Login generates the authentication token
// if the password was matched with the specified account.
func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		fmt.Println("err:", err)
		return
	}

	SendResponse(c, nil, model.SayHello{Message: "Welcome", UserName: d.Username})
}