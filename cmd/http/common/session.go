package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const (
	domain        = "localhost:8080"
	sessionKey    = "user_id"
	sessionMaxAge = 86400
)

func SetSession(c *gin.Context, userID int64) error {
	c.SetCookie(sessionKey, cast.ToString(userID), sessionMaxAge, "/", domain, true, true)
	return nil
}

func RemoveSession(c *gin.Context) error {
	c.SetCookie(sessionKey, "", sessionMaxAge, "/", domain, true, true)
	return nil
}

func GetSession(c *gin.Context) (int64, error) {
	userID, err := c.Cookie(sessionKey)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64(userID), nil
}
