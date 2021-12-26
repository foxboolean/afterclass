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
	return ua.us.UserInfo(nil, request)
}

type Result struct {
	code string
	msg string
	data interface{}
}

func NewResult(data interface{}, err error) *Result {
	if err!=nil {
		return &Result{
			code: "500",
			msg: err.Error(),
			data: nil,
		}
	}
	return &Result{
		code: "200",
		data: data,
	}
}

type RequestInfo struct {
	Id string
}

func NewRequestInfo(id string) *RequestInfo {
	return &RequestInfo{
		Id: id,
	}
}