package api

import (
	"log"
	"nitic-pbl-2022-01/pbl-back/src/auth"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"nitic-pbl-2022-01/pbl-back/src/response"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	type reqBody struct {
		Email     string
		Password  string
		IsStudent bool
	}

	var body reqBody
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	email, err := domain.ConstructEmail(body.Email)
	if err != nil {
		handleError(c, err)
		return
	}

	user := domain.ConstructUser(email, body.Password, body.IsStudent)

	_, err = impl.User.Create(user)
	if err != nil {
		handleError(c, err)
		return
	}

	token, err := auth.Signup(body.Email, body.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, token)
}

func Signin(c *gin.Context) {
	type reqBody struct {
		Email    string
		Password string
	}

	var body reqBody
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	token, err := auth.SignIn(body.Email, body.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, token)
}

func FetchSelf(c *gin.Context) {
	email, err := getEmail(c)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := repository.User.GetByEmail(email)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, response.ConvertUser(user))
}
