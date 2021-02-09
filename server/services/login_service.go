package services

import (
	"chatRoom/server/common/validate"
	"chatRoom/server/model"
	"chatRoom/server/repositories"
	"errors"
	"strings"
)

var LoginService = newLoginService()

func newLoginService() *loginService {
	return &loginService{}
}

type loginService struct {
}

func (s *loginService) SignUp(username, password, rePassword string) (*model.User, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	rePassword = strings.TrimSpace(rePassword)

	if len(username) == 0 {
		return nil, errors.New("昵称不能为空！")
	}

	if err := validate.IsPassword(password, rePassword); err != nil {
		return nil, err
	}

	if len(username) > 0 {
		if err := validate.IsUsername(username); err != nil {
			return nil, err
		}
		if s.existUserName(username) {
			return nil, errors.New("该用户名已注册！")
		}
	}

	user := &model.User{
		Username:    username,
		Password:    password,
		Avatar:      "",
		Description: "",
	}

	if err := repositories.UserRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *loginService) existUserName(username string) bool {
	return s.getUserByUserName(username) == nil
}

func (s *loginService) getUserByUserName(username string) *model.User {
	return repositories.UserRepository.FindOneByUserName(username)
}
