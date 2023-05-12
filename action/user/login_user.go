package user

import (
	"context"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/middleware"
	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	uRepo mysql.MySQL
}

func (u *UserLogin) Handle(ctx context.Context, req *entity.Users) (res entity.Users, err error) {
	user, err := u.uRepo.GetUserByEmail(ctx, req.Email)
	if err != nil || user == (entity.Users{}) {
		return
	}

	errComparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if errComparePass != nil {
		err = errComparePass
		return
	}

	token, errCreateToken := middleware.CreateToken(int(user.ID))
	if err != nil {
		err = errCreateToken
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
