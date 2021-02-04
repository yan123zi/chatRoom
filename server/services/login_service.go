package services

import (
	"chatRoom/server/model"
)

var LoginService=newLoginService()

func newLoginService() *loginService {
	return &loginService{}
}

type loginService struct {

}
func (s *loginService) SignUp(username,password,rePassword string) (*model.User,error) {

	return nil,nil

}
