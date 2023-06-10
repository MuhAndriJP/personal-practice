package user

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/MuhAndriJP/personal-practice.git/middleware"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
)

type UserLogin struct {
	uRepo mysql.MySQL
}

func (u *UserLogin) Handle(ctx context.Context, req *entity.Users) (res entity.Users, err error) {
	user, err := u.uRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	if user == (entity.Users{}) {
		err = errors.New(helper.StatusMessage[helper.NotFound])
		return
	}

	err = helper.CheckPassword(user.Password, req.Password)
	if err != nil {
		err = errors.New(helper.StatusMessage[helper.Unauthorized])
		return
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return
	}

	user.Token = token

	err = u.uRepo.Upsert(ctx, &user)
	if err != nil {
		return
	}

	res = user

	return
}

func NewUserLogin() *UserLogin {
	return &UserLogin{
		uRepo: mysql.NewSQL(),
	}
}
