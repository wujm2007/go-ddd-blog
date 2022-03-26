package user

import (
	"github.com/gin-gonic/gin"

	"go-ddd-blog/cmd/http/common"
	"go-ddd-blog/internal/service/user/api"
)

type handler struct {
	app *api.Application
}

func NewHandler(app *api.Application) *handler {
	return &handler{app: app}
}

func (h *handler) Signup(c *gin.Context) {
	req := struct {
		Name     string
		Password string
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	userID, err := h.app.Create(c, req.Name, req.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if err = common.SetSession(c, userID); err != nil {
		_ = c.Error(err)
		return
	}
	common.Success(c, gin.H{
		"user_id":   userID,
		"user_name": req.Name,
	})
}

func (h *handler) Login(c *gin.Context) {
	req := struct {
		Name     string
		Password string
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	userDTO, err := h.app.GetByNamePassword(c, req.Name, req.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if err = common.SetSession(c, userDTO.ID); err != nil {
		_ = c.Error(err)
		return
	}
	common.Success(c, gin.H{
		"user_id":   userDTO.ID,
		"user_name": userDTO.Name,
	})
}

func (h *handler) Logout(c *gin.Context) {
	if err := common.RemoveSession(c); err != nil {
		return
	}
	common.Success(c, gin.H{})
}
