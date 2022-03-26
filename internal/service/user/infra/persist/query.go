package persist

import (
	"context"

	"gorm.io/gorm"

	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/user/api/model"
	"go-ddd-blog/internal/service/user/api/query"
)

type userQuery struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) query.Query {
	return &userQuery{db: db}
}

func (q userQuery) GetByNamePassword(ctx context.Context, name, password string) (user *model.UserDTO, err error) {
	var po UserPO
	if err = q.db.WithContext(ctx).First(&po, "name = ?", name).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	if po.Password != password {
		return nil, errors.ErrForbidden.WithMessage("incorrect password")
	}

	return &model.UserDTO{
		ID:        po.ID,
		Name:      po.Name,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}, nil
}

func (q userQuery) Get(ctx context.Context, id int64) (user *model.UserDTO, err error) {
	var po UserPO
	if err = q.db.WithContext(ctx).First(&po, id).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	return &model.UserDTO{
		ID:        po.ID,
		Name:      po.Name,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}, nil
}
