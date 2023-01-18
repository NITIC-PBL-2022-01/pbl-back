package api

import (
	"log"
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
