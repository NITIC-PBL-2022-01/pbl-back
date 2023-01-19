package api

import (
	"errors"
	"log"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	e := c.Error(err).SetType(gin.ErrorTypePublic)
	if e != nil {
		log.Println(e)
	}
}

func strEpochToTime(epoch string) (time.Time, error) {
	n, err := strconv.Atoi(epoch)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(int64(n), 0), nil
}

func fold[T any, S any](array []T, fn func(a T) S) []S {
	newArray := []S{}
	for _, v := range array {
		newArray = append(newArray, fn(v))
	}

	return newArray
}

func getEmail(c *gin.Context) (domain.Email, error) {
	email, exist := c.Get("email")
	if !exist {
		c.JSON(401, map[string]string{"message": "Unauthorized"})
		return domain.Email(""), errors.New("Unauthorized")
	}

	return email.(domain.Email), nil
}
