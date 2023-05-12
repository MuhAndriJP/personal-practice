package user

import (
	"context"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
)

type GetUserByEmail struct {
	mySql mysql.MySQL
}

func (u *GetUserByEmail) Handle(ctx context.Context, req *entity.Users) (res entity.Users, err error) {
	user, err := u.mySql.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	res = user

	return
}

func NewGetUserByEmail() *GetUserByEmail {
	return &GetUserByEmail{}
}
