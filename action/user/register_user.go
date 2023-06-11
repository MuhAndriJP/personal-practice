package user

import (
	"context"
	"errors"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	uRepo mysql.IUser
}

func (u *UserRegister) Handle(ctx context.Context, req entity.Users) (err error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		err = errors.New(helper.StatusMessage[helper.InvalidArgument])
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	req.Password = string(hashedPassword)

	err = u.uRepo.Insert(ctx, &req)
	if err != nil {
		return
	}

	return
}

func NewUserRegister() *UserRegister {
	return &UserRegister{
		uRepo: mysql.NewSQL(),
	}
}
