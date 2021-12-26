package service

import (
	"afterclass/week_four/internal/biz"
	"afterclass/week_four/internal/dao"
	"context"
)

type UserService struct {
	ubiz *biz.UserBiz
}

func NewUserService(ubiz *biz.UserBiz) *UserService{
	return &UserService{
		ubiz: ubiz,
	}
}

func (u *UserService) UserInfo(ctx context.Context, id string) (*dao.User, error) {
	return u.ubiz.GetUserById(id)
}