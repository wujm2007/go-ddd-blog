package post

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"go-ddd-blog/cmd/http/common"
	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/post/api"
)

type handler struct {
	app *api.Application
}

func NewHandler(app *api.Application) *handler {
	return &handler{app: app}
}

func (h *handler) ListPost(c *gin.Context) {
	userID, ok := c.GetQuery("user_id")
	if !ok {
		_ = c.Error(errors.ErrBadRequest.WithMessage("invalid user_id"))
		return
	}

	posts, err := h.app.MGetByUserID(c, cast.ToInt64(userID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.Success(c, gin.H{"posts": posts})
}

func (h *handler) GetPost(c *gin.Context) {
	postID, ok := c.GetQuery("post_id")
	if !ok {
		_ = c.Error(errors.ErrBadRequest.WithMessage("invalid post_id"))
		return
	}

	post, err := h.app.Get(c, cast.ToInt64(postID), true)
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.Success(c, gin.H{"post": post})
}

func (h *handler) CreatPost(c *gin.Context) {
	userID, err := common.GetSession(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req := struct {
		Title   string
		Content string
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	postID, err := h.app.Create(c, userID, req.Title, req.Content)
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.Success(c, gin.H{"post_id": postID})
}

func (h *handler) ModifyPost(c *gin.Context) {
	userID, err := common.GetSession(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req := struct {
		PostID  int64  `json:"post_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	if err := h.app.Update(c, userID, req.PostID, req.Title, req.Content); err != nil {
		_ = c.Error(err)
		return
	}
	common.Success(c, gin.H{})
}

func (h *handler) Comment(c *gin.Context) {
	userID, err := common.GetSession(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req := struct {
		PostID  int64  `json:"post_id" binding:"gte=1"`
		Content string `json:"content" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	if err := h.app.AddComment(c, req.PostID, userID, req.Content); err != nil {
		_ = c.Error(err)
		return
	}
	common.Success(c, gin.H{})
}
