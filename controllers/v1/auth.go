package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/daystram/go-gin-gorm-boilerplate/datatransfers"
	"github.com/daystram/go-gin-gorm-boilerplate/handlers"
)

func POSTLogin(c *gin.Context) {
	var err error
	var user datatransfers.UserLogin
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	var token string
	if token, err = handlers.Handler.AuthenticateUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "incorrect username or password"})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: fmt.Sprintf("Bearer %s", token)})
}

func POSTRegister(c *gin.Context) {
	var err error
	var user datatransfers.UserSignup
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.RegisterUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Error: "failed registering user"})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Data: "user created"})
}
