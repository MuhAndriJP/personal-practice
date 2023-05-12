package mysql

import (
	"context"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"gorm.io/gorm"
)

type MySQL interface {
	Insert(context.Context, *entity.Users) error
	GetUserByEmail(context.Context, string) (entity.Users, error)
	Upsert(context.Context, *entity.Users) error
}

type SQL struct {
}

func (m *SQL) Insert(ctx context.Context, user *entity.Users) (err error) {
	if err = DB.Debug().Create(&user).Error; err != nil {
		return
	}
	return
}

func (m *SQL) GetUserByEmail(ctx context.Context, email string) (user entity.Users, err error) {
	user = entity.Users{}
	if err = DB.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return
	}
	return
}

func (m *SQL) Upsert(ctx context.Context, user *entity.Users) (err error) {
	if err = DB.Debug().Save(&user).Error; err != nil {
		return
	}
	return
}

func NewSQL() MySQL {
	return &SQL{}
}
