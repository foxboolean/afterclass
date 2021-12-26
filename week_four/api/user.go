package api

import "afterclass/week_four/internal/service"

type UserApi struct {
	us *service.UserService
}

func NewUserApi(userService *service.UserService) *UserApi {
	return &UserApi{
		us: userService,
	}
}

func (ua *UserApi) QueryUserInfo(request *RequestInfo) *Result {
	user, err := ua.us.UserInfo(nil, request.Id)
	if err != nil {
		return NewResult(nil, err)
	}
	return NewResult(user, err)
}