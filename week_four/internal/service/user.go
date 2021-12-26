package service

import (
	"afterclass/week_four/api"
	"afterclass/week_four/internal/biz"
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

func (u *UserService) UserInfo(ctx context.Context, request *api.RequestInfo) (*api.Result) {
	user, err := u.ubiz.GetUserById(request.Id)
	if err != nil {
		return api.NewResult(nil, err)
	}
	return api.NewResult(user, err)
}