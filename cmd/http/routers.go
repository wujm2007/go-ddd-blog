package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"go-ddd-blog/cmd/http/common"
	"go-ddd-blog/cmd/http/handler/post"
	"go-ddd-blog/cmd/http/handler/user"
	postApp "go-ddd-blog/internal/service/post/api"
	userApp "go-ddd-blog/internal/service/user/api"
)

func ErrorHandle(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	lastErr := c.Errors[len(c.Errors)-1]
	log.Default().Printf("error: %+v", lastErr)
	common.Error(c, lastErr)
}

func Init(userApp *userApp.Application, postApp *postApp.Application) *gin.Engine {
	r := gin.Default()
	r.Use(ErrorHandle)

	userV1 := r.Group("/api/user/v1")
	{
		h := user.NewHandler(userApp)
		userV1.POST("/signup", h.Signup)
		userV1.POST("/login", h.Login)
		userV1.POST("/logout", h.Logout)
	}

	postV1 := r.Group("/api/post/v1")
	{
		h := post.NewHandler(postApp)
		postV1.GET("/list", h.ListPost)
		postV1.GET("/get", h.GetPost)
		postV1.POST("/create", h.CreatPost)
		postV1.POST("/modify", h.ModifyPost)
		postV1.POST("/comment", h.Comment)
	}

	return r
}
