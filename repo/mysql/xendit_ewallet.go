package mysql

import (
	"context"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
)

type IEWallet interface {
	Insert(context.Context, *entity.EWalletPayment) error
}

type ewallet struct {
}

func (m *ewallet) Insert(ctx context.Context, callback *entity.EWalletPayment) (err error) {
	if err = DB.Debug().Create(&callback).Error; err != nil {
		return
	}
	return
}

func NewEWallet() IEWallet {
	return &ewallet{}
}
