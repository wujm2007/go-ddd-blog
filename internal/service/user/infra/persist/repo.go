package persist

import (
	"context"

	"gorm.io/gorm"

	"go-ddd-blog/internal/errors"
	"go-ddd-blog/internal/service/user/domain"
	"go-ddd-blog/internal/service/user/domain/aggregate/user"
)

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) domain.Repository {
	return &userRepo{db: db}
}

func (r userRepo) Save(ctx context.Context, user user.User) (id int64, err error) {
	po := marshalUser(user)
	db := r.db.WithContext(ctx)
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(po).Error; err != nil {
			return err
		}
		messagePOs := make([]MessagePO, len(user.Messages()))
		for i, message := range user.Messages() {
			messagePOs[i] = *marshalMessage(message)
		}
		if len(messagePOs) != 0 {
			return tx.Save(messagePOs).Error
		}
		return nil
	})
	return po.ID, errors.TransformGormErr(err)
}

func (r userRepo) Get(ctx context.Context, id int64) (user *user.User, err error) {
	var po UserPO
	if err = r.db.WithContext(ctx).First(&po, id).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	var messagePOs []MessagePO
	if err = r.db.WithContext(ctx).Find(&messagePOs, "user_id = ?", po.ID).Error; err != nil {
		return nil, errors.TransformGormErr(err)
	}

	return unmarshalUser(po, messagePOs), nil
}
