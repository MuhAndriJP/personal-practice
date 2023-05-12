package user

import (
	"context"
	"log"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	uRepo mysql.MySQL
}

func (u *UserRegister) Handle(ctx context.Context, req entity.Users) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	req.Password = string(hashedPassword)

	err = u.uRepo.Insert(ctx, &req)
	if err != nil {
		log.Println("err")
		return
	}

	return
}

func NewUserRegister() *UserRegister {
	return &UserRegister{
		uRepo: mysql.NewSQL(),
	}
}
