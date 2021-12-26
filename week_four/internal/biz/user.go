package biz

import (
	"afterclass/week_four/internal/dao"
	"errors"
)

type UserBiz struct {
	dao *dao.UserDAO
}

func NewUserBiz(dao *dao.UserDAO) *UserBiz {
	return &UserBiz{
		dao: dao,
	}
}

func (u UserBiz) GetUserById(id string) (*dao.User, error) {
	if id == "" {
		return nil, errors.New("empty userid")
	}
	return u.dao.GetUser(id)
}