package services

import (
	"chatRoom/server/common/validate"
	"chatRoom/server/model"
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

	if err := validate.IsUsername(username); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *loginService) existUserName(username string) bool {

}

func (s *loginService) getUserByUserName() {

}
