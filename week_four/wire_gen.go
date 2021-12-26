// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"afterclass/week_four/internal/biz"
	"afterclass/week_four/internal/dao"
	"afterclass/week_four/internal/service"
)

// Injectors from wire.go:

func InitUserService() *service.UserService {
	db := dao.NewDB()
	userDAO := dao.NewUserDAO(db)
	userBiz := biz.NewUserBiz(userDAO)
	userService := service.NewUserService(userBiz)
	return userService
}
